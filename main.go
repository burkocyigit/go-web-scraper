package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
	"sync"


	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/mbndr/figlet4go"
)

type post struct {
	title string
	description string
	date string
}

func main() {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	renderStr, _ := ascii.RenderOpts("YAVUZLAR", options)
	fmt.Print(renderStr)
	
	ascii2 := figlet4go.NewAsciiRender()
	renderStr2, _ := ascii2.Render("Web Scraper Tool")
	fmt.Print(renderStr2)

	datePtr := flag.Bool("date", false, "exclude dates")
	descPtr := flag.Bool("description", false, "exclude descriptions")
	hackerNewsPtr := flag.Bool("1", false, "hackernews.com posts")
	secondNewsPtr := flag.Bool("2", false, "cybersecuritynews.com posts")
	thirdNewsPtr := flag.Bool("3", false, "cyware.com posts")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage of %s: \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	
	printNews(*hackerNewsPtr, *datePtr, *descPtr)
	printNews(*secondNewsPtr, *datePtr, *descPtr)
	printNews(*thirdNewsPtr, *datePtr, *descPtr)
}

func getTitlesFromHackerNews(date bool, desc bool) ([10]post, error) {
	url := "https://thehackernews.com/"

	var posts = []post{}

	res, err := http.Get(url)

	if err != nil {
		return [10]post(posts), err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return [10]post(posts), errors.New(strconv.FormatInt(int64(res.StatusCode), 10))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return [10]post(posts), err
	}

	var postTitles, postDescs, postDates []string
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		doc.Find(".home-right").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find(".home-title").Text()
			postTitles = append(postTitles, title)
		})
	}()

	if !date {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doc.Find(".item-label").Each(func(i int, selection *goquery.Selection) {
				date := selection.Find(".h-datetime").Text()
				postDates = append(postDates, date)
			})
		}()
	} else {
		for i := 0; i < 10; i++ {
			postDates = append(postDates, "")
		}
	}

	if !desc {
		wg.Add(1)
		go func(){
			defer wg.Done()

			doc.Find(".home-right").Each(func(i int, selection *goquery.Selection) {
				desc := selection.Find(".home-desc").Text()
				postDescs = append(postDescs, desc)
			})
		}() 
	} else {
		for i := 0; i < 10; i++ {
			postDescs = append(postDescs, "")
		}
	}
		
	wg.Wait()

	for i := 0; i < 10; i++ {
		var p post = post{title: postTitles[i], description: postDescs[i], date: postDates[i]}
		posts = append(posts, p) 
	}	
		
	return [10]post(posts), nil
}

func getTitlesFromCybersecurityNews(date bool, desc bool) ([10]post, error) {
	url := "https://cybersecuritynews.com/"

	var posts = []post{}

	res, err := http.Get(url)

	if err != nil {
		return [10]post(posts), err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return [10]post(posts), errors.New(strconv.FormatInt(int64(res.StatusCode), 10))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return [10]post(posts), err
	}

	var postTitles, postDescs, postDates []string
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		doc.Find(".entry-title").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("a").Text()
			postTitles = append(postTitles, title)
		})
	}()

	if !date {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doc.Find(".td-post-date").Each(func(i int, selection *goquery.Selection) {
				date := selection.Find("time").Text()
				postDates = append(postDates, date)
			})
		}()
	} else {
		for i := 0; i < 10; i++ {
			postDates = append(postDates, "")
		}
	}

	if !desc {
		wg.Add(1)
		go func(){
			defer wg.Done()

			doc.Find(".item-details").Each(func(i int, selection *goquery.Selection) {
				desc := selection.Find(".td-excerpt").Text()
				postDescs = append(postDescs, desc)
			})
		}() 
	} else {
		for i := 0; i < 10; i++ {
			postDescs = append(postDescs, "")
		}
	}
		
	wg.Wait()

	for i := 0; i < 10; i++ {
		var p post = post{title: postTitles[i], description: postDescs[i], date: postDates[i]}
		posts = append(posts, p) 
	}	
		
	return [10]post(posts), nil
}

func getTitlesFromCyware(date bool, desc bool) ([10]post, error) {
	url := "https://cyware.com/cyber-security-news-articles"

	var posts = []post{}

	res, err := http.Get(url)

	if err != nil {
		return [10]post(posts), err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return [10]post(posts), errors.New(strconv.FormatInt(int64(res.StatusCode), 10))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return [10]post(posts), err
	}

	var postTitles, postDescs, postDates []string
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		doc.Find(".cy-panel__body").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("a h1").Text()
			postTitles = append(postTitles, title)
		})
	}()

	if !date {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doc.Find(".cy-panel__body").Each(func(i int, selection *goquery.Selection) {
				date := selection.Find("span").Text()
				postDates = append(postDates, date)
			})
		}()
	} else {
		for i := 0; i < 10; i++ {
			postDates = append(postDates, "")
		}
	}

	if !desc {
		wg.Add(1)
		go func(){
			defer wg.Done()

			doc.Find(".cy-panel__body").Each(func(i int, selection *goquery.Selection) {
				desc := selection.Find(".cy-card__description").Text()
				postDescs = append(postDescs, desc)
			})
		}() 
	} else {
		for i := 0; i < 10; i++ {
			postDescs = append(postDescs, "")
		}
	}
		
	wg.Wait()

	for i := 0; i < 10; i++ {
		var p post = post{title: postTitles[i], description: postDescs[i], date: postDates[i]}
		posts = append(posts, p) 
	}	
		
	return [10]post(posts), nil
}

func printNews(hackerNewsPtr bool, datePtr bool, descPtr bool) {
	if !hackerNewsPtr {
		return
	} else {
		posts, err := getTitlesFromHackerNews(datePtr, descPtr)

		if err != nil {
			log.Fatal(err)
		}
	
		for i, v := range posts {
			color.Red("News %v:", i+1)
			fmt.Printf("%v\n", v.title)
			if !descPtr {
				color.Blue("Description:")
				fmt.Printf("%v\n", v.description)
			}

			if !datePtr {
				color.Magenta("Date:")
				fmt.Printf("%v\n\n\n", v.date)
			}
		}
	}
}