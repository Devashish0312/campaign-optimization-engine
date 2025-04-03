package predictive

import (
	"testing"
)

func TestPredictiveModel(t *testing.T) {
	model := &PredictiveModel{}
	model.Train()

	predictedCPC := model.Predict()

	if predictedCPC <= 0 || predictedCPC > 0.5 {
		t.Errorf("Expected predicted CPC to be in range [0, 0.5], got %v", predictedCPC)
	}
}
