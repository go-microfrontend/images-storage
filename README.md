# Бекенд-воркер обработки изображений

## Описание

Go-сервис для обработки информации о товарах. Его цель заключается в том, чтобы предоставить presigned-ссылки на изображения из S3-хранилища или кеша. Это производится путём запросов на исполнения задач, которые воркер вытаскивает из очереди Temporal.

## Зависимости

Для запуска экземпляра воркера необходимо, чтобы был установлен Docker и работал Temporal-кластер, представленный в [orchestration](https://github.com/go-microfrontend/orchestration) репозитории.

## Запуск

Для запуска необходимо:
```
make build
make up
```
