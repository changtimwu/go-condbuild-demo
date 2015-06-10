TARGETS=main.xcat main.linkstreet

main.xcat: main.go hal_xcat.go xcat
	go build -tags xcat -o $@

main.linkstreet: main.go hal_linkstreet.go linkstreet
	go build -tags linkstreet -o $@


clean:
	rm -f $(TARGETS)
