package telegraph_test

import (
	"log"

	"gitlab.com/toby3d/telegraph"
)

// Content in a string format (for this example).
// Be sure to wrap every media in a <figure> tag, okay? Be easy.
const data = `
    <figure>
        <img src="/file/6a5b15e7eb4d7329ca7af.jpg"/>
    </figure>
    <p><i>Hello</i>, my name is <b>Page</b>, <u>look at me</u>!</p>
    <figure>
        <iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
        <figcaption>
            Yes, you can embed youtube, vimeo and twitter widgets too!
        </figcaption>
    </figure>
`

var (
	account *telegraph.Account //nolint:gochecknoglobals
	page    *telegraph.Page    //nolint:gochecknoglobals
	content []telegraph.Node   //nolint:gochecknoglobals
)

func errCheck(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Example_fastStart() {
	var err error
	// Create new Telegraph account.
	requisites := telegraph.Account{
		ShortName: "toby3d", // required

		// Author name/link can be epmty. So secure. Much anonymously. Wow.
		AuthorName: "Maxim Lebedev",       // optional
		AuthorURL:  "https://t.me/toby3d", // optional
	}
	account, err = telegraph.CreateAccount(requisites)
	errCheck(err)

	// Make sure that you have saved acc.AuthToken for create new pages or make
	// other actions by this account in next time!

	// Format content to []telegraph.Node array. Input data can be string, []byte
	// or io.Reader.
	content, err = telegraph.ContentFormat(data)
	errCheck(err)

	// Boom!.. And your text will be understandable for Telegraph. MAGIC.

	// Create new Telegraph page
	pageData := telegraph.Page{
		Title:   "My super-awesome page", // required
		Content: content,                 // required

		// Not necessarily, but, hey, it's just an example.
		AuthorName: account.AuthorName, // optional
		AuthorURL:  account.AuthorURL,  // optional
	}
	page, err = account.CreatePage(pageData, false)
	errCheck(err)

	// Show link from response on created page.
	log.Println("Kaboom! Page created, look what happened:", page.URL)
}

func ExampleCreateAccount() {
	var err error
	account, err = telegraph.CreateAccount(telegraph.Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	})
	errCheck(err)

	log.Println("AccessToken:", account.AccessToken)
	log.Println("AuthURL:", account.AuthorURL)
	log.Println("ShortName:", account.ShortName)
	log.Println("AuthorName:", account.AuthorName)
}

func ExampleAccount_EditAccountInfo() {
	var err error
	account, err = account.EditAccountInfo(telegraph.Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	})
	errCheck(err)

	log.Println("AuthURL:", account.AuthorURL)
	log.Println("ShortName:", account.ShortName)
	log.Println("AuthorName:", account.AuthorName)
}

func ExampleAccount_GetAccountInfo() {
	info, err := account.GetAccountInfo(
		telegraph.FieldShortName,
		telegraph.FieldPageCount,
	)
	errCheck(err)

	log.Println("ShortName:", info.ShortName)
	log.Println("PageCount:", info.PageCount, "pages")
}

func ExampleAccount_RevokeAccessToken() {
	var err error
	// You must rewrite current variable with account structure for further usage.
	account, err = account.RevokeAccessToken()
	errCheck(err)

	log.Println("AccessToken:", account.AccessToken)
}

func ExampleAccount_CreatePage() {
	var err error
	page, err = account.CreatePage(telegraph.Page{
		Title:      "Sample Page",
		AuthorName: account.AuthorName,
		Content:    content,
	}, true)
	errCheck(err)

	log.Println(page.Title, "by", page.AuthorName, "has been created!")
	log.Println("PageURL:", page.URL)
}

func ExampleAccount_EditPage() {
	var err error
	page, err = account.EditPage(telegraph.Page{
		Title:      "Sample Page",
		AuthorName: account.AuthorName,
		Content:    content,
	}, true)
	errCheck(err)

	log.Println("Page on", page.Path, "path has been updated!")
	log.Println("PageURL:", page.URL)
}

func ExampleGetPage() {
	info, err := telegraph.GetPage("Sample-Page-12-15", true)
	errCheck(err)

	log.Println("Getted info about", info.Path, "page:")
	log.Println("Author:", info.AuthorName)
	log.Println("Views:", info.Views)
	log.Println("CanEdit:", info.CanEdit)
}

func ExampleAccount_GetPageList() {
	list, err := account.GetPageList(0, 3)
	errCheck(err)

	log.Println("Getted", list.TotalCount, "pages")
	for i := range list.Pages {
		p := list.Pages[i]
		log.Printf("%s: %s\n~ %s\n\n", p.Title, p.URL, p.Description)
	}
}

func ExampleGetViews() {
	pagePath := "Sample-Page-12-15"
	views, err := telegraph.GetViews(pagePath, 2016, 12)
	errCheck(err)

	log.Println(pagePath, "has been viewed", views.Views, "times")
}

func ExampleContentFormat() {
	const data = `<figure>
<img src="http://telegra.ph/file/6a5b15e7eb4d7329ca7af.jpg" /></figure>
<p><i>Hello</i>, my name is <b>Page</b>, <u>look at me</u>!</p>
<figure><iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
<figcaption>Yes, you can embed youtube, vimeo and twitter widgets too!</figcaption>
</figure>`

	var err error
	content, err = telegraph.ContentFormat(data)
	errCheck(err)

	log.Printf("Content: %#v", content)
}
