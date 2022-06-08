# Repository Backend FGD Kelompok 40

## Git Branching Convention
### Branching Rules
#### Master Branch
Merge to master branch only for deployment.
#### Development Branch
Merge to development branch after a feature is created. Pull request/commit pushed in development branch should trigger github actions continuous integration workflow to test the code.
#### Other Branches
To develop new features, create new "feature/feature-name" branch for each feature.

## Project Structure
We are using golang-standards' project layout along with bxcodec's clean architecture:
Golang standards project layout:
https://github.com/golang-standards/project-layout
Bxcodec's clean architecture:
https://github.com/bxcodec/go-clean-arch