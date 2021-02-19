## Settings git hooks

Используется [pre-commit](https://pre-commit.com/#install)

Установка pre-commit hook скриптов:

    pre-commit install --hook-type pre-commit --hook-type commit-msg

Запуск конкретного pre-commit hook:

    pre-commit run <hook_id>

Запуск pre-commit hooks для всех файлов проекта:

    pre-commit run --all-files