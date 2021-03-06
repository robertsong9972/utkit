package config

const IncreasePrepareScript = `
if ! command -v gocov; then
  echo "Install gocov"
  go get github.com/axw/gocov/gocov
fi
if ! command -v gocov-xml; then
	echo "Install gocov-xml"
	go get github.com/AlekSi/gocov-xml
fi
if [ ! -d localfiles ];then
   mkdir -p localfiles
fi
go mod tidy
`

const RunTest = `
go test -cover ./... -gcflags=all=-l -coverprofile=localfiles/cover.out
`

const GenerateXml = `
gocov convert localfiles/cover.out | gocov-xml > localfiles/coverage.xml
`

const IncreaseDiffCalScript = `
diff-cover localfiles/coverage.xml --compare-branch=origin/master --html-report localfiles/report.html
`

const MakeGitIgnoreFile = `
echo "localfiles/" >> .gitignore
`
