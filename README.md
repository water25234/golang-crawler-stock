# golang-crawler-stock


// Find the review items
	// doc.Find(".hasBorder .odd").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the title
	// 	title := s.Find("pre").Text()
	// 	if len(title) > 0 {
	// 		fmt.Printf("%s\n", title)
	// 	}

	// })

	// // Find the review items
	// doc.Find(".hasBorder .odd").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the title
	// 	title := s.Find("pre").Text()
	// 	fmt.Printf("Review %d: %s\n", i, title)
	// })

	// jar, _ := cookiejar.New(nil)

	// // Instantiate default collector
	// c := colly.NewCollector(
	// 	colly.AllowedDomains("jdata.yuanta.com.tw"),
	// )

	// c.WithTransport(&http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// })

	// //setup our client based on the cookies data
	// c.SetCookieJar(jar)

	// q, _ := queue.New(
	// 	1, // Number of consumer threads
	// 	&queue.InMemoryQueueStorage{MaxSize: 100000}, // Use default queue storage
	// )

	// q.AddURL("https://jdata.yuanta.com.tw/z/zc/zcr/zcr0.djhtm?b=Q&a=2330")

	// var result string

	// c.OnHTML("body", func(e *colly.HTMLElement) {

	// 	e.DOM.Find("table").Eq(2).Find("td").Each(func(i int, s *goquery.Selection) {
	// 		if i > 1 {
	// 			if i >= 2 && i <= 10 {
	// 				result = result + s.Text() + "\n"
	// 			}

	// 			if i >= 47 && i <= 55 {
	// 				result = result + s.Text() + "\n"
	// 			}

	// 			if i >= 147 && i <= 155 {
	// 				result = result + s.Text() + "\n"
	// 			}

	// 		}
	// 	})

	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// q.Run(c)

	// fmt.Println(result)