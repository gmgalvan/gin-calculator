package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSum(t *testing.T) {

	tests := []struct {
		name          string
		inputOperands []byte
		wantResult    Result
		code          int
	}{

		{
			name:          "Correct sum",
			inputOperands: []byte(`{"quantities": [1,2,3]}`),
			wantResult: Result{
				Result: 6,
			},
			code: 200,
		},
		{
			name:          "Error during unmarshall",
			inputOperands: []byte(`bad json`),
			wantResult:    Result{},
			code:          400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)

			r := gin.Default()

			r.POST("/sum", Sum)
			req, err := http.NewRequest(http.MethodPost, "/sum", bytes.NewBuffer(tt.inputOperands))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.code {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.code, w.Code)
			}

			if tt.wantResult.Result != 0 {
				var bodyResult Result
				err = json.Unmarshal(w.Body.Bytes(), &bodyResult)
				if err != nil {
					t.Error("error during unmarshalling")
				}
				if bodyResult.Result != tt.wantResult.Result {
					t.Errorf("expected body does not match error = %v, wantErr %v", &bodyResult, tt.wantResult)
				}
			}

		})
	}
}
