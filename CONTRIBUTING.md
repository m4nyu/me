# Contributing Guidelines

## Core Principles

- **Clean Code First**: Prioritize readability, maintainability, and simplicity
- **Atomic Everything**: Small, focused changes in commits, functions, and PRs
- **No Attribution**: Commits must not include tool or AI attribution
- **Language Agnostic**: Standards apply regardless of tech stack

---

## Git Standards

### Commit Convention

All commits must follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

#### Commit Types

- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code change without fixing bug or adding feature
- `perf`: Performance improvement
- `test`: Add/update tests
- `docs`: Documentation only
- `style`: Formatting (no code change)
- `chore`: Build process or auxiliary tool changes
- `revert`: Revert a previous commit
- `build`: Changes affecting build system
- `ci`: CI configuration changes

#### Commit Message Rules

- Subject: 50 characters max, imperative mood
- Body: 72 characters per line, explain what and why
- Footer: Breaking changes, issue references
- **No tool attribution** (no "Generated with", "Co-Authored-By: AI", etc.)

### Atomic Commits

Each commit must:
- Contain a single logical change
- Build successfully
- Pass all tests
- Be revertable independently
- Change fewer than 200 lines (when possible)

**Example:**
```bash
# Good: Selective staging
git add -p
git commit -m "fix(auth): validate email format"

# Bad: Bulk staging with vague message
git add .
git commit -m "fix stuff"
```

---

## Code Quality

### DRY Principle

Avoid duplication by extracting common patterns:

```go
// BAD: Duplication
func validateEmail(email string) bool {
    pattern := "email_regex_pattern"
    return match(email, pattern)
}

func checkEmail(email string) bool {
    pattern := "email_regex_pattern"
    return match(email, pattern)
}

// GOOD: Single source of truth
const EmailPattern = "email_regex_pattern"

func validateEmail(email string) bool {
    return match(email, EmailPattern)
}
```

### AHA Principle (Avoid Hasty Abstractions)

- Allow duplication **twice** before abstracting
- Prefer duplication over wrong abstraction
- Refactor when pattern becomes clear

### Code Metrics

| Metric | Target | Maximum |
|--------|--------|---------|
| Cyclomatic Complexity | ≤ 5 | 10 |
| Cognitive Complexity | ≤ 10 | 15 |
| Function Length | ≤ 20 lines | 50 lines |
| File Length | ≤ 200 lines | 400 lines |
| Code Coverage | ≥ 80% | - |

### Naming Conventions

Use descriptive, full-word names:

```go
// GOOD
userAuthenticationToken := generateToken()
func calculateTotalPrice() int
type UserAccountManager struct

// BAD
tkn := genTkn()
func calc() int
type UAM struct
```

### Self-Documenting Code

Code should explain itself without comments:

```go
// BAD: Requires comments
func proc(d []int) []int {
    // Process data for validation
    return filter(d, func(x int) bool { return x > 0 })
}

// GOOD: Self-explanatory
func filterPositiveNumbers(numbers []int) []int {
    return filter(numbers, isPositive)
}
```

---

## Development Workflow

### Before Starting

1. Check existing code to avoid duplication
2. Review README.md for project setup
3. Test all documented commands work
4. Plan approach before coding

### Development Process

```bash
# 1. Create feature branch
git checkout -b feat/feature-name

# 2. Make atomic changes with frequent commits
# ... develop ...

# 3. Run quality checks
make lint
make test
make build

# 4. Update documentation if needed
# ... update README.md ...

# 5. Clean up commit history
git rebase -i main
```

### Quality Checklist

Before submitting changes:

- [ ] All tests pass
- [ ] No linter errors
- [ ] Code coverage ≥ 80%
- [ ] Build succeeds
- [ ] Documentation updated
- [ ] Atomic commits with clear messages

---

## Testing Standards

### Test Structure

```go
// Arrange-Act-Assert pattern
func TestCalculateTotal(t *testing.T) {
    // Arrange
    input := createTestData()

    // Act
    result := calculateTotal(input)

    // Assert
    assert.Equal(t, expected, result)
}
```

### Coverage Requirements

- Minimum: 80% overall
- Critical paths: 100%
- New code: 90%
- UI components: 70%

---

## Pull Request Guidelines

### PR Size

- **Ideal**: < 200 lines
- **Maximum**: 400 lines
- **If larger**: Split into multiple PRs

### PR Template

```markdown
## What
Brief description of changes

## Why
Context and motivation

## How to Test
1. Step-by-step instructions
2. Expected behavior

## Checklist
- [ ] Tests pass
- [ ] Documentation updated
- [ ] No linter warnings
- [ ] Self-reviewed
```

### Review Focus

1. Logic correctness
2. Edge case handling
3. Performance implications
4. Security considerations
5. Code clarity

---

## Security Requirements

All code must:

- [ ] Contain no hardcoded secrets
- [ ] Validate all inputs
- [ ] Prevent SQL injection
- [ ] Protect against XSS
- [ ] Implement CSRF protection
- [ ] Use up-to-date dependencies

---

## Performance Standards

- Response time: < 200ms (p95)
- Memory usage: Within defined limits
- CPU usage: < 80% under normal load
- Startup time: < 5 seconds

---

## Documentation Standards

### Code Comments

Use comments only when necessary:

1. Complex algorithms requiring explanation
2. Non-obvious business logic
3. Workarounds with context
4. **Never** for obvious code

### Function Documentation

```go
// ValidateEmail checks if email conforms to RFC 5322 standard.
// Returns true if valid, false otherwise.
// Panics if email is not a string type.
func ValidateEmail(email string) bool {
    // Implementation
}
```

---

## Decision Framework

### When to Abstract

1. **Third occurrence** of similar code
2. **Clear pattern** emerges
3. **Multiple files** need same logic

### When to Split Files

1. File exceeds **200 lines**
2. Multiple **responsibilities**
3. Different **rate of change**

### When to Refactor

1. Before adding new feature to messy code
2. When fixing bugs in complex area
3. When code violates standards
4. **Not** during critical deadlines

---

## Project-Specific Commands

```bash
# Development
make dev          # Start development server with hot reload
make build        # Build the application
make run          # Build and run

# Quality
make lint         # Check code style
make test         # Run tests
make clean        # Clean build artifacts

# Deployment
make deploy-prod    # Deploy to production
make deploy-staging # Deploy to staging
```

---

## Remember

1. **Small commits** - One logical change per commit
2. **Clear messages** - Explain why, not what
3. **Test everything** - Especially edge cases
4. **Document decisions** - Use ADRs when needed
5. **No attribution** - Clean, professional commits only
6. **Measure quality** - Meet coverage and complexity targets
