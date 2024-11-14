package ftracker

import (
    "fmt"
    "math"
)

// Основные константы, необходимые для расчетов.
const (
    lenStep   = 0.65  // средняя длина шага (в метрах)
    mInKm     = 1000  // количество метров в километре.
    minInH    = 60    // количество минут в часе.
    kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
    cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
    return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    distance := distance(action)
    return distance / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
    switch {
	case trainingType == "Бег":
		distance := distance(action) // дистанция в километрах
		speed := meanSpeed(action, duration) // скорость при беге в км/ч
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := distance(action) 
		speed := meanSpeed(action, duration) 
		calories := WalkingSpentCalories(action, duration, weight, height) 
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := distance(action)
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
    runSpeedIndex  = 18   // множитель средней скорости.
    aveCalories = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
    speed := meanSpeed(action, duration) // средняя скорость
    res := ((runSpeedIndex * speed * aveCalories) * weight / mInKm * duration * minInH)
    return res
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
    walkWeightIndex = 0.035 // множитель массы тела.
    heightIndex = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
    speed := meanSpeed(action, duration) * kmhInMsec // средняя скорость в м/с
    res := ((walkWeightIndex * weight + (math.Pow(speed, 2) / height * cmInM) * heightIndex * weight) * duration * minInH)
    return res
}

// Константы для расчета калорий, расходуемых при плавании.
const (
    swimCals   = 1.1  // среднее количество сжигаемых колорий при плавании относительно скорости.
    swimWeightIndex = 2    // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
    speed := swimmingMeanSpeed(lengthPool, countPool, duration) // скорость при плавании в км/ч
    res := (speed + swimCals) * 2 * weight * duration
    return res
}
