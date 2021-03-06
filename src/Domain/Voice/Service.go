package Voice

import (
	"io"
	"math/cmplx"
	"math"
	"mime/multipart"

	"github.com/mjibson/go-dsp/fft"
	"github.com/youpy/go-wav"
)

type ServiceInterface interface {
	CalculatePowerSpectrum(multipart.File) []float64
	CosSimilarity(sample []float64, training []float64) float64
	Add(name string, powerSpectrum []float64) (*Voice, error)
	Get(id string) (*Voice, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) ServiceInterface {
	service := new(Service)
	service.repository = repository
	return service
}

func (s Service) Add(name string, powerSpectrum []float64) (*Voice, error) {
	voice := new(Voice)
	voice.Name = name
	voice.PowerSpectrum = powerSpectrum
	return s.repository.Add(voice)
}

func (s Service) Get(id string) (*Voice, error) {
	return s.repository.Get(id)
}

func (s Service) CalculatePowerSpectrum(file multipart.File) []float64 {
	// Wavファイルの読み込み 
	reader := wav.NewReader(file)

	defer file.Close()

	signals := make([]float64, 0)

	// 解析はモノラルチャンネルで行うので片方のチャンネルだけ取得
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			signals = append(signals, reader.FloatValue(sample, 0))
		}
	}

	// 高速フーリエ変換
	fftResult := fft.FFTReal(signals)

	// フーリエ変換の結果は複素数なので絶対値を取得し, パワースペクトルに変換する
	powerSpectrum := make([]float64, 0)
	for _, r := range fftResult {
		powerSpectrum = append(powerSpectrum, cmplx.Abs(r))
	}

	// 人間の音声の周波数帯は100~2000hzなのでその範囲で取得する
	// 後々の配列 -> 文字列変換が非常に遅くなるため
	return powerSpectrum[LOW_FREQUENCY:HIGH_FREQUENCY]
}

func (s Service) CosSimilarity(sample []float64, training []float64) float64 {
	sSize := math.Sqrt(calculateDot(sample, sample))
	tSize := math.Sqrt(calculateDot(training, training))
	return calculateDot(sample, training) / (sSize * tSize)
}

func minInt(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

func calculateDot(x []float64, y []float64) float64 {
	result := 0.0;
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result
}
