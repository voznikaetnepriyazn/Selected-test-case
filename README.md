# 📋 log-linter — Линтер для анализа лог-записей в Go

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org)
[![golangci-lint](https://img.shields.io/badge/golangci--lint-plugin-compatible-brightgreen)](https://golangci-lint.run)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> **log-linter** — плагин для `golangci-lint`, который анализирует лог-записи в коде на Go и проверяет их соответствие установленным правилам стиля и безопасности.

---

## 🎯 Назначение

Линтер помогает поддерживать единый стиль логирования в проекте и предотвращает утечку чувствительных данных через логи.

---

## ✅ Проверяемые правила

### 1️⃣ Сообщение должно начинаться со строчной буквы

| Статус | Пример |
|--------|---------|
| ❌ | `slog.Info("Starting server")` |
| ✅ | `slog.Info("starting server")` |

**Обоснование**: единообразие стиля, соответствие конвенциям многих лог-фреймворков.

---

### 2️⃣ Сообщение должно быть на английском языке

| Статус | Пример |
|--------|---------|
| ❌ | `log.Info("ошибка подключения")` |
| ✅ | `log.Info("connection failed")` |

**Обоснование**: международная совместимость, упрощение поиска в логах, единый стандарт команды.

---

### 3️⃣ Запрещены спецсимволы и эмодзи

| Статус | Пример |
|--------|---------|
| ❌ | `slog.Info("started! 🚀")`<br>`log.Warn("warning...")` |
| ✅ | `slog.Info("started")`<br>`log.Warn("warning")` |

**Запрещено**:
- Эмодзи и символы за пределами ASCII (`\u{1F600}`–`\u{1F64F}` и др.)
- Повторяющаяся пунктуация: `!!!`, `???`, `...` (3+ символа подряд)

**Разрешено**: базовая пунктуация `. , ! ? : ; " ' - ( ) [ ] { } / @ # $ % & * + = < > _ | \ ` ~`

**Обоснование**: чистота логов, совместимость с парсерами, избежание проблем с кодировками.

---

### 4️⃣ Запрещены потенциально чувствительные данные

| Статус | Пример |
|--------|---------|
| ❌ | `log.Info("password: " + pwd)`<br>`slog.Debug("api_key=" + key)` |
| ✅ | `log.Info("user authenticated")`<br>`slog.Debug("api request completed")` |

**Проверяемые паттерны** (регистронезависимые):
- password, passwd, pwd
- api_key, api-key, apikey
- token, secret, auth
- credentials, cred
- private_key, private-key

**Важно**: правило срабатывает **только при конкатенации строк** (`"msg" + var`), чтобы не блокировать безопасные сообщения вида `"password field is required"`.

**Обоснование**: предотвращение утечек секретов в логи, соответствие стандартам безопасности (OWASP, PCI DSS).

---

## ⚙️ Технические требования

### Язык и зависимости

| Требование | Значение |
|-----------|----------|
| Минимальная версия Go | **1.22+** |
| Основной пакет анализа | `golang.org/x/tools/go/analysis` |
| Проверка языка | `golang.org/x/text` (опционально) |
| Тестирование | `github.com/stretchr/testify` |

### Совместимость

| Компонент | Поддержка |
|-----------|-----------|
| **golangci-lint** | ✅ Плагин через `.custom-gcl.yml` |
| **go vet** | ✅ Standalone режим через `-vettool` |
| **CI/CD** | ✅ GitHub Actions, GitLab CI |

### Поддерживаемые логгеры

| Пакет | Методы |
|-------|--------|
| `log` | Print, Printf, Println, Fatal, Panic |
| `log/slog` | Info, Debug, Warn, Error, Log, LogAttrs |
| `go.uber.org/zap` (Logger) | Info, Debug, Warn, Error, Panic, Fatal |
| `go.uber.org/zap` (SugaredLogger) | Info, Debug, Warn, Error, Panic, Fatal |

---

## 🚀 Установка и использование

### Вариант 1: Через golangci-lint (рекомендуется)

**1. Создайте файл `.custom-gcl.yml`:**

```yaml
version: v1.55.2
name: custom-gcl-loglinter
destination: ./bin

plugins:
  - module: 'github.com/yourname/log-linter'
    import: 'github.com/yourname/log-linter/pkg/loglinter'
    path: .
