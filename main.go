package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// Input struct for JSON decoding
type Input struct {
	Values []float64 `json:"values"`
}

func calculateTask1(I, t, Sm, jEk float64) string {
	U := 10.0

	Im := Sm / (2 * (math.Sqrt(3.0) * U))
	ImPA := 2 * Im

	// Розрахунок економічного перерізу та оптимального перерізу кабеля
	sEk := Im / jEk
	s := I * 1000 * math.Sqrt(t) / 92

	output := fmt.Sprintf(
		"Розрахунковий струм для нормального режиму: %.0f A\n"+
			"Розрахунковий струм для післяаварійного режиму: %.0f A\n"+
			"Економічний переріз: %.0f мм^2\n"+
			"Оптимальний переріз: %.0f мм^2\n",
		math.Round(Im),
		math.Round(ImPA),
		math.Round(sEk),
		math.Round(s),
	)

	return output
}

func calculateTask2(Usn float64) string {
	// P - Потужність трансформатора, кВт
	// SNom - Номінальна потужність трансформатора, МВА
	P := 200.0
	SNom := 6.3

	// Конвертація введених користувачем даних

	// Розрахунок опорів елементів заступної схеми та сумарного опору
	Xc := (Usn * Usn) / P
	Xt := (Usn * Usn * Usn) / SNom / 100
	X := Xc + Xt

	// Розрахунок початкового діючого значення струму трифазного КЗ, А
	Ip0 := Usn / math.Sqrt(3.0) / X

	// Формування вихідного рядка
	output := fmt.Sprintf(
		"Опори елементів заступної схеми: %.2f Ом і %.2f Ом\nСумарний опір: %.2f Ом\nПочаткове діюче значення струму трифазного КЗ: %.2f А\n",
		Xc, Xt, X, Ip0,
	)

	return output
}

func calculateTask3() string {
	SNom := 6.3
	Uh := 115.0
	Ul := 11.0
	UkMax := 11.1
	Rcn := 10.65
	Xcn := 24.02
	RcMin := 34.88
	XCMin := 65.68

	XT := UkMax * math.Pow(Uh, 2.0) / 100 / SNom
	Xsh := Xcn + XT
	Zsh := math.Sqrt(math.Pow(Rcn, 2.0) + math.Pow(Xsh, 2.0))

	XshMin := XCMin + XT
	ZshMin := math.Sqrt(math.Pow(RcMin, 2.0) + math.Pow(XshMin, 2.0))

	Ish3 := Uh * 1000 / math.Sqrt(3.0) / Zsh
	Ish2 := Ish3 * math.Sqrt(3.0) / 2

	IshMin3 := Uh * 1000 / math.Sqrt(3.0) / ZshMin
	IshMin2 := IshMin3 * math.Sqrt(3.0) / 2

	Kpr := math.Pow(Ul, 2.0) / math.Pow(Uh, 2.0)

	RshN := Rcn * Kpr
	XshN := Xsh * Kpr
	ZshN := math.Sqrt(math.Pow(RshN, 2.0) + math.Pow(XshN, 2.0))

	RshMinN := RcMin * Kpr
	XshMinN := XshMin * Kpr
	ZshMinN := math.Sqrt(math.Pow(RshMinN, 2.0) + math.Pow(XshMinN, 2.0))

	IshN3 := Ul * 1000 / math.Sqrt(3.0) / ZshN
	IshN2 := IshN3 * math.Sqrt(3.0) / 2

	IshMinN3 := Ul * 1000 / math.Sqrt(3.0) / ZshMinN
	IshMinN2 := IshMinN3 * math.Sqrt(3.0) / 2

	Rl := 7.91
	Xl := 4.49

	RSumN := Rl + RshN
	XSumN := Xl + XshN
	ZSumN := math.Sqrt(math.Pow(RSumN, 2.0) + math.Pow(XSumN, 2.0))

	RSumMinN := Rl + RshMinN
	XSumMinN := Xl + XshMinN
	ZSumMinN := math.Sqrt(math.Pow(RSumMinN, 2.0) + math.Pow(XSumMinN, 2.0))

	Iln3 := Ul * 1000 / math.Sqrt(3.0) / ZSumN
	Iln2 := Iln3 * math.Sqrt(3.0) / 2

	IlMinN3 := Ul * 1000 / math.Sqrt(3.0) / ZSumMinN
	IlMinN2 := IlMinN3 * math.Sqrt(3.0) / 2

	output := fmt.Sprintf(`
Опір на шинах в нормальному режимі: %.2f Ом
Опір на шинах в мінімальному режимі: %.2f Ом
Сила трифазного струму на шинах в нормальному режимі: %.2f А
Сила двофазного струму на шинах в нормальному режимі: %.2f А
Сила трифазного струму на шинах в мінімальному режимі: %.2f А
Сила двофазного струму на шинах в мінімальному режимі: %.2f А
Коефіцієнт приведення для визначення дійсних струмів: %.2f
Опір на шинах в нормальному режимі: %.2f Ом
Опір на шинах в мінімальному режимі: %.2f Ом
Сила трифазного струму на шинах в нормальному режимі: %.2f А
Сила двофазного струму на шинах в нормальному режимі: %.2f А
Сила трифазного струму на шинах в мінімальному режимі: %.2f А
Сила двофазного струму на шинах в мінімальному режимі: %.2f А
Опір в точці 10 в нормальному режимі: %.2f Ом
Опір в точці 10 в мінімальному режимі: %.2f Ом
Сила трифазного струму в точці 10 в нормальному режимі: %.2f А
Сила двофазного струму в точці 10 в нормальному режимі: %.2f А
Сила трифазного струму в точці 10 в мінімальному режимі: %.2f А
Сила двофазного струму в точці 10 в мінімальному режимі: %.2f А
`, Zsh, ZshMin, Ish3, Ish2, IshMin3, IshMin2, Kpr, ZshN, ZshMinN, IshN3, IshN2, IshMinN3, IshMinN2, ZSumN, ZSumMinN, Iln3, Iln2, IlMinN3, IlMinN2)

	return output
}

func calculator1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(input.Values) != 4 {
		http.Error(w, "Invalid number of inputs", http.StatusBadRequest)
		return
	}
	result := calculateTask1(input.Values[0], input.Values[1], input.Values[2], input.Values[3])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func calculator2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(input.Values) != 1 {
		http.Error(w, "Invalid number of inputs", http.StatusBadRequest)
		return
	}
	result := calculateTask2(input.Values[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func calculator3Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(input.Values) != 0 {
		http.Error(w, "Invalid number of inputs", http.StatusBadRequest)
		return
	}
	result := calculateTask3()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/calculator1", calculator1Handler)
	http.HandleFunc("/api/calculator2", calculator2Handler)
	http.HandleFunc("/api/calculator3", calculator3Handler)

	fmt.Println("Server running at http://localhost:8084")
	http.ListenAndServe(":8084", nil)
}
