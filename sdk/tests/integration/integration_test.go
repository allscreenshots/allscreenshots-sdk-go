// Package integration contains integration tests for the Allscreenshots SDK.
//
// These tests require a valid API key set in the ALLSCREENSHOTS_API_KEY environment variable.
// Run with: go test -v ./tests/integration/... -tags=integration
package integration

import (
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/allscreenshots/allscreenshots-sdk-go/pkg/allscreenshots"
)

// TestResult represents the result of a single test case.
type TestResult struct {
	ID           string
	Name         string
	URL          string
	Device       string
	FullPage     bool
	Passed       bool
	ErrorMessage string
	ImageData    string // Base64 encoded
	ExecutionMs  int64
}

// TestReport represents the full test report.
type TestReport struct {
	SDKName        string
	Language       string
	Version        string
	Timestamp      string
	TotalTests     int
	PassedTests    int
	FailedTests    int
	TotalTimeMs    int64
	Results        []TestResult
	OSInfo         string
	RuntimeVersion string
}

var testCases = []struct {
	ID       string
	Name     string
	URL      string
	Device   string
	FullPage bool
	WantErr  bool
}{
	{"IT-001", "Basic Desktop Screenshot", "https://github.com", "Desktop HD", false, false},
	{"IT-002", "Basic Mobile Screenshot", "https://github.com", "iPhone 14", false, false},
	{"IT-003", "Basic Tablet Screenshot", "https://github.com", "iPad", false, false},
	{"IT-004", "Full Page Desktop", "https://github.com", "Desktop HD", true, false},
	{"IT-005", "Full Page Mobile", "https://github.com", "iPhone 14", true, false},
	{"IT-006", "Complex Page", "https://github.com/anthropics/claude-code", "Desktop HD", false, false},
	{"IT-007", "Invalid URL", "not-a-valid-url", "Desktop HD", false, true},
	{"IT-008", "Unreachable URL", "https://this-domain-does-not-exist-12345.com", "Desktop HD", false, true},
}

func TestIntegration(t *testing.T) {
	apiKey := os.Getenv("ALLSCREENSHOTS_API_KEY")
	if apiKey == "" {
		t.Fatal("ALLSCREENSHOTS_API_KEY environment variable is not set")
	}

	client := allscreenshots.NewClient(
		allscreenshots.WithAPIKey(apiKey),
		allscreenshots.WithTimeout(120*time.Second),
	)

	report := TestReport{
		SDKName:        "allscreenshots-sdk-go",
		Language:       "Go",
		Version:        "1.0.0",
		Timestamp:      time.Now().Format(time.RFC3339),
		OSInfo:         fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		RuntimeVersion: runtime.Version(),
	}

	var totalTime int64

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := TestResult{
				ID:       tc.ID,
				Name:     tc.Name,
				URL:      tc.URL,
				Device:   tc.Device,
				FullPage: tc.FullPage,
			}

			start := time.Now()

			ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
			defer cancel()

			req := &allscreenshots.ScreenshotRequest{
				URL:      tc.URL,
				Device:   tc.Device,
				FullPage: tc.FullPage,
			}

			imageData, err := client.Screenshot(ctx, req)
			result.ExecutionMs = time.Since(start).Milliseconds()
			totalTime += result.ExecutionMs

			if tc.WantErr {
				if err != nil {
					result.Passed = true
					result.ErrorMessage = fmt.Sprintf("Expected error: %v", err)
					t.Logf("Test %s passed: expected error occurred - %v", tc.ID, err)
				} else {
					result.Passed = false
					result.ErrorMessage = "Expected error but got success"
					t.Errorf("Test %s failed: expected error but got success", tc.ID)
				}
			} else {
				if err != nil {
					result.Passed = false
					result.ErrorMessage = err.Error()
					t.Errorf("Test %s failed: %v", tc.ID, err)
				} else if len(imageData) == 0 {
					result.Passed = false
					result.ErrorMessage = "Empty image data returned"
					t.Errorf("Test %s failed: empty image data", tc.ID)
				} else {
					result.Passed = true
					result.ImageData = base64.StdEncoding.EncodeToString(imageData)
					t.Logf("Test %s passed: received %d bytes", tc.ID, len(imageData))
				}
			}

			report.Results = append(report.Results, result)
		})
	}

	report.TotalTests = len(testCases)
	report.TotalTimeMs = totalTime

	for _, r := range report.Results {
		if r.Passed {
			report.PassedTests++
		} else {
			report.FailedTests++
		}
	}

	// Generate HTML report
	if err := generateHTMLReport(&report); err != nil {
		t.Errorf("Failed to generate HTML report: %v", err)
	}
}

func generateHTMLReport(report *TestReport) error {
	const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Integration Test Report - {{.SDKName}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
            background: #f5f5f5;
            color: #333;
            line-height: 1.6;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        header {
            background: #1a1a1a;
            color: white;
            padding: 30px 20px;
            margin-bottom: 30px;
        }
        header h1 {
            font-size: 24px;
            font-weight: 500;
            margin-bottom: 10px;
        }
        .meta {
            display: flex;
            gap: 30px;
            font-size: 14px;
            color: #888;
        }
        .summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin-bottom: 30px;
        }
        .summary-card {
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
        }
        .summary-card .label {
            font-size: 12px;
            text-transform: uppercase;
            color: #666;
            margin-bottom: 5px;
        }
        .summary-card .value {
            font-size: 32px;
            font-weight: 600;
        }
        .summary-card.passed .value { color: #22c55e; }
        .summary-card.failed .value { color: #ef4444; }
        .test-results {
            background: white;
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .test-results h2 {
            padding: 20px;
            font-size: 18px;
            border-bottom: 1px solid #eee;
        }
        .test-item {
            border-bottom: 1px solid #eee;
            padding: 20px;
        }
        .test-item:last-child {
            border-bottom: none;
        }
        .test-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 15px;
        }
        .test-id {
            font-family: monospace;
            background: #f0f0f0;
            padding: 2px 8px;
            border-radius: 4px;
            font-size: 12px;
            margin-right: 10px;
        }
        .test-name {
            font-weight: 500;
        }
        .badge {
            padding: 4px 12px;
            border-radius: 4px;
            font-size: 12px;
            font-weight: 500;
        }
        .badge.passed {
            background: #dcfce7;
            color: #166534;
        }
        .badge.failed {
            background: #fee2e2;
            color: #991b1b;
        }
        .test-details {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
            gap: 10px;
            font-size: 14px;
            color: #666;
            margin-bottom: 15px;
        }
        .test-details span {
            background: #f5f5f5;
            padding: 5px 10px;
            border-radius: 4px;
        }
        .test-image {
            margin-top: 15px;
        }
        .test-image img {
            max-width: 100%;
            max-height: 400px;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .error-message {
            background: #fee2e2;
            color: #991b1b;
            padding: 10px 15px;
            border-radius: 4px;
            font-size: 14px;
            margin-top: 10px;
        }
        footer {
            margin-top: 30px;
            padding: 20px;
            text-align: center;
            color: #666;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <h1>Integration Test Report</h1>
            <div class="meta">
                <span>SDK: {{.SDKName}}</span>
                <span>Version: {{.Version}}</span>
                <span>Run: {{.Timestamp}}</span>
            </div>
        </div>
    </header>

    <div class="container">
        <div class="summary">
            <div class="summary-card">
                <div class="label">Total Tests</div>
                <div class="value">{{.TotalTests}}</div>
            </div>
            <div class="summary-card passed">
                <div class="label">Passed</div>
                <div class="value">{{.PassedTests}}</div>
            </div>
            <div class="summary-card failed">
                <div class="label">Failed</div>
                <div class="value">{{.FailedTests}}</div>
            </div>
            <div class="summary-card">
                <div class="label">Total Time</div>
                <div class="value">{{.TotalTimeMs}}ms</div>
            </div>
        </div>

        <div class="test-results">
            <h2>Test Results</h2>
            {{range .Results}}
            <div class="test-item">
                <div class="test-header">
                    <div>
                        <span class="test-id">{{.ID}}</span>
                        <span class="test-name">{{.Name}}</span>
                    </div>
                    <span class="badge {{if .Passed}}passed{{else}}failed{{end}}">
                        {{if .Passed}}PASSED{{else}}FAILED{{end}}
                    </span>
                </div>
                <div class="test-details">
                    <span>URL: {{.URL}}</span>
                    <span>Device: {{.Device}}</span>
                    <span>Full Page: {{.FullPage}}</span>
                    <span>Time: {{.ExecutionMs}}ms</span>
                </div>
                {{if and .Passed .ImageData}}
                <div class="test-image">
                    <img src="data:image/png;base64,{{.ImageData}}" alt="Screenshot">
                </div>
                {{end}}
                {{if and (not .Passed) .ErrorMessage}}
                <div class="error-message">{{.ErrorMessage}}</div>
                {{end}}
                {{if and .Passed .ErrorMessage}}
                <div class="test-details">
                    <span>{{.ErrorMessage}}</span>
                </div>
                {{end}}
            </div>
            {{end}}
        </div>
    </div>

    <footer>
        <div class="container">
            Environment: {{.OSInfo}} | Runtime: {{.RuntimeVersion}}
        </div>
    </footer>
</body>
</html>`

	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create("test-report.html")
	if err != nil {
		return fmt.Errorf("failed to create report file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, report); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
