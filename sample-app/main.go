package main
import (
 "context"
 "encoding/json"
 "fmt"
 "io"
 "net/http"
 "net/url"
 "os"
 "strings"
 "time"
 sdk "github.com/allscreenshots/allscreenshots-sdk-go/v2"
)
func env(name,fallback string) string { if v:=os.Getenv(name);v!="" {return v}; return fallback }
func main() { if err:=run();err!=nil {fmt.Fprintln(os.Stderr,"SDK request failed. Check authentication, quota, and the API response.");os.Exit(1)} }
func run() error {
 config:=sdk.NewConfiguration(); config.Servers=sdk.ServerConfigurations{{URL:env("ALLSCREENSHOTS_BASE_URL","https://api.allscreenshots.com")}};config.HTTPClient=&http.Client{Timeout:180*time.Second}
 client:=sdk.NewAPIClient(config);ctx:=context.WithValue(context.Background(),sdk.ContextAPIKeys,map[string]sdk.APIKey{"ApiKey":{Key:os.Getenv("ALLSCREENSHOTS_API_KEY")}})
 mode:="quota";if len(os.Args)>1 {mode=os.Args[1]}
 if mode=="quota" {q,_,e:=client.UsageAPI.GetQuota(ctx).Execute();if e!=nil{return e};return json.NewEncoder(os.Stdout).Encode(q)}
 request:=sdk.NewScreenshotRequest(env("ALLSCREENSHOTS_URL","https://example.com"));request.SetResponseType("URL");request.SetFormat("png")
 key:=env("ALLSCREENSHOTS_IDEMPOTENCY_KEY",fmt.Sprintf("demo-%d",time.Now().UnixNano()))
 var file *os.File
 if mode=="sync" {
  raw,_,e:=client.ScreenshotAPI.CaptureSync(ctx).ScreenshotRequest(*request).IdempotencyKey(key).Execute();if e!=nil{return e}
  var metadata sdk.ScreenshotJsonResponse; e=json.NewDecoder(raw).Decode(&metadata);raw.Close();os.Remove(raw.Name());if e!=nil{return e}
  parsed,e:=url.Parse(metadata.GetResultUrl());if e!=nil{return e};parts:=strings.Split(parsed.Path,"/")
  file,_,e=client.ScreenshotAPI.GetSyncCaptureResult(ctx,parts[len(parts)-2]).Execute();if e!=nil{return e}
 } else if mode=="async" {
  job,_,e:=client.JobAPI.CreateAsyncJob(ctx).ScreenshotRequest(*request).IdempotencyKey(key).Execute();if e!=nil{return e};deadline:=time.Now().Add(180*time.Second)
  for {status,_,e:=client.JobAPI.GetJobStatus(ctx,job.GetId()).Execute();if e!=nil{return e};state:=status.GetStatus();if state=="COMPLETED"{break};if state=="FAILED"||state=="CANCELLED"{return fmt.Errorf("capture failed")};if time.Now().After(deadline){return fmt.Errorf("polling timed out")};time.Sleep(2*time.Second)}
  file,_,e=client.JobAPI.GetJobResult(ctx,job.GetId()).Execute();if e!=nil{return e}
 } else {return fmt.Errorf("expected quota, sync, or async")}
 defer os.Remove(file.Name());defer file.Close();output:=env("ALLSCREENSHOTS_OUTPUT","capture.png");out,e:=os.Create(output);if e!=nil{return e};_,e=io.Copy(out,file);out.Close();if e!=nil{return e};return json.NewEncoder(os.Stdout).Encode(map[string]string{"output":output})
}
