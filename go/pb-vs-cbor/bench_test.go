package bench

import (
	"encoding/json"
	"testing"

	pb "github.com/edgelaboratories/pb-vs-cbor/pb"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

type request struct {
	Run        *run     `json:"run"`
	Asset      string   `json:"asset"`
	Scenarios  []uint32 `json:"scenarios"`
	Currencies []string `json:"currencies"`
}

type run struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	Keyspace string `json:"keyspace"`
}

type response struct {
	Status    string             `json:"status"`
	Error     string             `json:"error"`
	Timestamp string             `json:"pricedAt"`
	Asset     string             `json:"asset"`
	Currency  string             `json:"currency"`
	Run       *run               `json:"run"`
	Results   map[string]float64 `json:"results"`
}

func BenchmarkRequestProtobuf(b *testing.B) {
	baseRequest := pb.Request{
		Run: &pb.Run{
			Id:       1,
			Date:     "2019-06-08",
			Keyspace: "risk_run",
		},
		Asset:      "cc6a7c4f-8b51-4518-93ba-288974798485",
		Scenarios:  generateScenarios(),
		Currencies: []string{"local", "USD", "CAD", "EUR", "GBP", "CHF"},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Read
		data, err := proto.Marshal(&baseRequest)
		assert.NoError(b, err)
		// Write
		var req pb.Request
		assert.NoError(b, proto.Unmarshal(data, &req))
	}
}

func BenchmarkRequestCBOR(b *testing.B) {
	baseRequest := request{
		Run: &run{
			Id:       1,
			Date:     "2019-06-08",
			Keyspace: "risk_run",
		},
		Asset:      "cc6a7c4f-8b51-4518-93ba-288974798485",
		Scenarios:  generateScenarios(),
		Currencies: []string{"local", "USD", "CAD", "EUR", "GBP", "CHF"},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Read
		data, err := json.Marshal(&baseRequest)
		assert.NoError(b, err)
		// Write
		var req request
		assert.NoError(b, json.Unmarshal(data, &req))
		_, err = encodeToHexCBOR(req)
		assert.NoError(b, err)
	}
}

func generateScenarios() []uint32 {
	output := []uint32{}
	for i := 0; i < 5000; i++ {
		output = append(output, uint32(i))
	}
	return output
}

func BenchmarkResponseProtobuf(b *testing.B) {
	baseResponse := pb.Response{
		Run: &pb.Run{
			Id:       1,
			Date:     "2019-06-08",
			Keyspace: "risk_run",
		},
		Asset:     "cc6a7c4f-8b51-4518-93ba-288974798485",
		Results:   generateResults(),
		Currency:  "local",
		Timestamp: "2019-06-09",
		Error:     "some error",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Read
		data, err := proto.Marshal(&baseResponse)
		assert.NoError(b, err)
		// Write
		var req pb.Response
		assert.NoError(b, proto.Unmarshal(data, &req))
	}
}

func BenchmarkResponseCBOR(b *testing.B) {
	baseResponse := response{
		Run: &run{
			Id:       1,
			Date:     "2019-06-08",
			Keyspace: "risk_run",
		},
		Asset:     "cc6a7c4f-8b51-4518-93ba-288974798485",
		Results:   generateResults(),
		Currency:  "local",
		Timestamp: "2019-06-09",
		Error:     "some error",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Read
		data, err := json.Marshal(&baseResponse)
		assert.NoError(b, err)
		// Write
		var req request
		assert.NoError(b, json.Unmarshal(data, &req))
		_, err = encodeToHexCBOR(req)
		assert.NoError(b, err)
	}
}

func generateResults() map[string]float64 {
	output := make(map[string]float64)
	for i := 0; i < 5000; i++ {
		output[string(i)] = float64(i)
	}
	return output
}
