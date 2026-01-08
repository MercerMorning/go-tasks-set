package main

// Задача: Принимаем запрос, делаем параллельно запросы к:
// - API A: https://api.service1.com/data
// - API B: https://api.service2.com/info
// Объединить результаты, вернуть JSON
// Важно: использовать context для отмены при таймауте
