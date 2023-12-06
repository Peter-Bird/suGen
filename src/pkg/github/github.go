package github

import "path/filepath"

type Github struct{}

func (g *Github) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, ".github", "workflows"),
		filepath.Join(path, ".github", "ISSUE_TEMPLATE"),
	}
}

func (g *Github) GetFiles(path, name string) map[string]string {
	return map[string]string{
		path + "/.github/PULL_REQUEST_TEMPLATE.md":          g.GenPullReqTemplate(),
		path + "/.github/CODEOWNERS":                        g.GenCodeOwners(),
		path + "/.github/FUNDING.yml":                       g.GenFunding(),
		path + "/.github/SECURITY.md":                       g.GenSecurity(),
		path + "/.github/ISSUE_TEMPLATE/bug_report.md":      g.GenBugReportTemplate(),
		path + "/.github/ISSUE_TEMPLATE/feature_request.md": g.GenFeatureReqTemplate(),
		path + "/.github/workflows/go.yml":                  g.GenWorkflow(),
	}
}

func (g *Github) GenPullReqTemplate() string {

	return `# Pull Request Template

## Description
Please include a summary of the change and which issue is fixed. Also include relevant motivation and context. List any dependencies that are required for this change.

Fixes # (issue number)

## Type of change
Please delete options that are not relevant.
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] This change requires a documentation update

## How Has This Been Tested?
Please describe the tests that you ran to verify your changes. Provide instructions so we can reproduce. Please also list any relevant details for your test configuration.
- [ ] Test A
- [ ] Test B

## Checklist:
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published in downstream modules
- [ ] I have checked my code and corrected any misspellings

## Additional Information
Add any other context or screenshots about the feature request here.
`
}

func (g *Github) GenCodeOwners() string {

	return `# CODEOWNERS file for a Golang project using Fyne and MongoDB
	# Owners are automatically requested for review on pull requests that modify code they own.
	
	# Default owner for everything in the repo
	*       @default-owner
	
	# Golang backend code
	/backend/                       @backend-team
	/backend/**/*.go                @go-developers
	
	# Fyne GUI code
	/gui/                           @frontend-team
	/gui/**/*.go                    @fyne-team
	
	# MongoDB related scripts and configurations
	/database/                      @database-team
	/**/*.bson                      @mongo-experts
	/**/*.mongo                     @mongo-experts
	
	# CI/CD Pipeline configurations
	/.github/workflows/             @devops-team
	
	# Docker configurations
	/Dockerfile                     @devops-team
	/docker-compose.yml             @devops-team
	
	# Documentation
	/docs/                          @documentation-team
	
	# Test files
	/**/*._test.go                  @quality-assurance
	
	# Specific file patterns
	/go.mod                         @dependency-managers
	/go.sum                         @dependency-managers
	
	# Assign a team to review changes in the project's README and LICENSE files
	/README.md                      @project-managers
	/LICENSE                        @legal-team
	`
}

func (g *Github) GenWorkflow() string {

	return `name: Go

	on:
	  push:
		branches: [ main ]
	  pull_request:
		branches: [ main ]
	
	jobs:
	
	  build:
		name: Build
		runs-on: ubuntu-latest
		steps:
	
		- name: Set up Git repository
		  uses: actions/checkout@v2
	
		- name: Set up Go
		  uses: actions/setup-go@v2
		  with:
			go-version: '1.16'  # The Go version to download (if necessary) and use.
	
		- name: Build
		  run: go build -v ./...
	
		- name: Run tests
		  run: go test -v ./...
	
	  # Additional job for MongoDB integration (if applicable)
	  mongodb-integration:
		name: MongoDB Integration Tests
		runs-on: ubuntu-latest
		services:
		  mongodb:
			image: mongo
			ports:
			  - 27017:27017
			options: --health-cmd 'mongo --eval "db.runCommand({ ping: 1 })" --health-interval 10s --health-timeout 5s --health-retries 5'
		steps:
	
		- name: Set up Git repository
		  uses: actions/checkout@v2
	
		- name: Set up Go
		  uses: actions/setup-go@v2
		  with:
			go-version: '1.16'
	
		- name: Run MongoDB integration tests
		  run: |
			export MONGODB_URI=mongodb://localhost:27017
			go test -v ./... -tags=mongodb
	`
}

func (g *Github) GenBugReportTemplate() string {

	return `---
	name: Bug report
	about: Create a report to help us improve
	title: ''
	labels: 'bug'
	assignees: ''
	
	---
	
	# Bug Report
	
	## Description
	A clear and concise description of what the bug is.
	
	## To Reproduce
	Steps to reproduce the behavior:
	1. Go to '...'
	2. Click on '....'
	3. Scroll down to '....'
	4. See error
	
	## Expected Behavior
	A clear and concise description of what you expected to happen.
	
	## Screenshots
	If applicable, add screenshots to help explain your problem.
	
	## Environment
	- OS: [e.g. iOS]
	- Browser [e.g. chrome, safari]
	- Version [e.g. 22]
	- Other relevant information [e.g. network conditions, specific configurations, etc.]
	
	## Additional Context
	Add any other context about the problem here. For example, the time when the problem occurred or any potential impact of the bug.
	
	## Possible Solution
	`
}

func (g *Github) GenFeatureReqTemplate() string {

	return `---
	name: Feature Request
	about: Suggest an idea for this project
	title: ''
	labels: 'enhancement'
	assignees: ''
	
	---
	
	# Feature Request
	
	## Is your feature request related to a problem? Please describe.
	A clear and concise description of what the problem is. Ex. I'm always frustrated when [...]
	
	## Describe the solution you'd like
	A clear and concise description of what you want to happen.
	
	## Describe alternatives you've considered
	A brief description of any alternative solutions or features you've considered.
	
	## Additional context
	Add any other context or screenshots about the feature request here. If you have specific examples or mockups, feel free to include them.
	
	## Potential Benefits
	Explain the benefits this feature would bring and why it would be a valuable addition to the project.
	
	## Possible Implementation
	If you have an idea of how this feature might be implemented, please outline it here. Your input can be very helpful to guide the development process.
	
	`
}

func (g *Github) GenFunding() string {

	return `# These are supported funding model platforms

github: [your-github-username]  # Replace with your GitHub Sponsor username

patreon: [your-patreon-username]  # Replace with your Patreon username

open_collective: [your-open-collective-username]  # Replace with your Open Collective username

ko_fi: [your-ko-fi-username]  # Replace with your Ko-fi username

tidelift: [your-tidelift-package-name]  # Replace with your Tidelift package name

community_bridge: [your-community-bridge-project-name]  # Replace with your Community Bridge project name

liberapay: [your-liberapay-username]  # Replace with your Liberapay username

issuehunt: [your-issuehunt-username]  # Replace with your IssueHunt username

otechie: [your-otechie-username]  # Replace with your Otechie username

custom: ['https://www.your-custom-funding-link.com']  # Replace with your custom funding link
`
}

func (g *Github) GenContributing() string {
	return `# Contributing to YourGoProject

We appreciate your interest in contributing to YourGoProject! Whether it's reporting a bug, discussing the current state of the code, submitting a fix, proposing new features, or becoming a maintainer, all contributions are welcome.

## We Develop with GitHub

We use GitHub to host code, track issues and feature requests, and accept pull requests.

## GitHub Flow

All code changes happen through pull requests. To contribute, follow these steps:

1. Fork the repo and create your branch from main.
2. Write clear, logical, and performant code. Please adhere to the [Go coding standards](https://golang.org/doc/effective_go.html).
3. Ensure any new code is accompanied by relevant tests.
4. If you've changed APIs, or added new ones, update the documentation accordingly.
5. Make sure the tests pass (go test ./...).
6. Lint your code (golint) and ensure it's formatted (gofmt -s).
7. Submit your pull request.

## Any Contributions You Make Will Be Under the Same License

When you submit code changes, your submissions are understood to be under the same [LICENSE](LICENSE) that covers the project.

## Report Bugs Using GitHub's Issues

Report a bug by [opening a new issue](https://github.com/yourusername/YourGoProject/issues/new); it's that easy!

### Write Detailed Bug Reports

A good bug report includes:

- A summary and/or background.
- Steps to reproduce. Be specific!
  - Give sample code if you can.
- What you expected to happen.
- What actually happens.
- Any additional information (e.g., why you think this might be happening, or things you tried that didn't work).

## Coding Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Use go fmt for formatting your code.
- Keep line lengths reasonable.
- Avoid package-level globals.

## License

By contributing, you agree that your contributions will be licensed under its MIT License.

## References

This document was adapted from the open-source contribution guidelines template.`
}

func (g *Github) GenSecurity() string {

	return `# Security Policy

	## Supported Versions
	
	Use this section to tell people about which versions of your project are currently being supported with security updates.
	
	| Version | Supported          |
	| ------- | ------------------ |
	| 1.2.x   | :white_check_mark: |
	| 1.1.x   | :white_check_mark: |
	| 1.0.x   | :x:                |
	| < 1.0   | :x:                |
	
	## Reporting a Vulnerability
	
	Your contributions and reports are sincerely appreciated. We encourage you to report vulnerabilities in a responsible manner.
	
	Please follow the steps below for reporting a security issue:
	
	1. **Do Not Open a GitHub Issue for the Vulnerability**: To avoid public disclosure, do not report the vulnerability directly through GitHub Issues.
	
	2. **Email Your Report**: Send a detailed report to [security@example.com](mailto:security@example.com). Include everything needed to reproduce the vulnerability. If you're not sure whether a vulnerability is real, you're still welcome to send a report.
	
	3. **Wait for Response**: We will review your report and respond as quickly as possible. We may request additional information or guidance on reproducing the issue.
	
	4. **Public Disclosure Timing**: After you've reported a vulnerability, we request that you do not publicly disclose it until we've had the chance to investigate and address the issue. We will work with you to ensure that an appropriate fix is applied and, when possible, will coordinate with you on releasing an announcement.
	
	5. **Credit**: After the vulnerability has been resolved, we will acknowledge your contributions in the project's release notes (unless you prefer to remain anonymous).
	
	## Security Related Configuration
	
	(If applicable, provide any specific details about security-related configurations for your Go project. This might include recommended configuration settings, features to enable/disable for enhanced security, etc.)
	
	## More Information
	
	(Any additional information about your security processes, future plans, contact information for security inquiries, etc.)
	
	Remember to replace '[security@example.com]' with the actual contact email for security reports, and ensure the version table accurately reflects your project's current state. The additional sections can be modified or omitted according to your project's needs.
	`
}
