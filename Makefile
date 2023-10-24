run:
	/Users/araz/go/bin/gin --appPort 8080 --all -i run main.go

git:
	git add .
	git commit -m "Build"
	git push -u origin main