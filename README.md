# linxdatacenter-test-task
## Задание
Есть 2 файла с данными о продуктах (наименование, цена, рейтинг) в 2-х форматах - CSV и JSON.
Необходимо написать программу, которая считывает данные из переданного в параметре файле, и
выводит «самый дорогой продукт» и «с самым высоким рейтингом».

Предусмотреть, что файлы могут быть огромными.

Репозиторий должен содержать Dockerfile для сборки готового приложения в docker среде.
## Запуск
Запуск:
    `go run cmd/main.go <path to file>`

Сборка image:

`docker build --force-rm -t linxdatacenter-test .`

Запуск контейнера:

`docker run --rm -v <PATH TO FILE>:/workdir/<FILE NAME> linxdatacenter-test <FILE NAME>`

### Пример

`docker run --rm -v $PWD/data/data.csv:/workdir/db.csv linxdatacenter-test db.csv`