package template

import "strings"

const (
	productPageSummaryReplaceTarget = "{SUMMARY}"
	productPageDetailsReplaceTarget = "{DETAILS}"
	productPageReviewsReplaceTarget = "{REVIEWS}"
)

var productPageHTML = strings.Join(
	[]string{`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Simple Bookstore App</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    margin: 0;
                    padding: 0;
                }
                header {
                    background-color: #f1f1f1;
                    padding: 20px;
                    text-align: center;
                }
                nav {
                    background-color: #222;
                    overflow: hidden;
                    padding: 14px 16px;
                }
                nav a {
                    color: #f1f1f1;
                    display: inline-block;
                    padding: 14px 16px;
                    text-align: center;
                    text-decoration: none;
                }
                nav button {
                    background-color: #337ab7;
                    border: none;
                    color: white;
                    float: right;
                    font-size: 14px;
                    margin-top: 8px;
                    padding: 5px 16px;
                }
                main {
                    display: flex;
                    justify-content: center;
                    padding: 20px;
                }
                .book-info {
                    max-width: 800px;
                }
                h3 {
                    text-align: center;
                    color: #337ab7;
                }
                p {
                    text-align: justify;
                }
                h4 {
                    text-align: center;
                    color: #337ab7;
                }
                .book-details,
                .book-reviews {
                    font-size: 18px;
                    padding: 10px;
                }
                .book-details dl {
                    display: grid;
                    grid-template-columns: auto 1fr;
                    grid-column-gap: 10px;
                    grid-row-gap: 4px;
                }
                .book-reviews blockquote {
                    margin-bottom: 20px;
                }
                .book-reviews small {
                    display: block;
                }
                .star {
                    color: gold;
                    font-size: 20px;
                    margin-right: 5px;
                }
                .empty-star {
                    color: #ddd;
                    font-size: 20px;
                    margin-right: 5px;
                }
            </style>
        </head>
        <body>
        <header>
            <h1>Bookinfo Product Page</h1>
        </header>
        <nav>
            <a href="#">BookInfo Sample</a>
            <button type="button">Sign in</button>
        </nav>
        <main>
            <div class="book-info">`,
		productPageSummaryReplaceTarget,
		"",
		`
            <div class="book-details">
                <h4>Book Details</h4>
        `,
		productPageDetailsReplaceTarget,
		`
            </div>
            <div class="book-reviews">
                <h4>Book Reviews</h4>
        `,
		productPageReviewsReplaceTarget,
		`
            </div>
        </div>
        </main>
        <!-- Add the following code right before the closing </body> tag -->
        <div id="login-modal" style="display: none; position: fixed; z-index: 1; left: 0; top: 0; width: 100%; height: 100%; overflow: auto; background-color: rgba(0, 0, 0, 0.4);">
        <div style="background-color: #fefefe; margin: 15% auto; padding: 20px; border: 1px solid #888; width: 30%;">
            <h2>Login</h2>
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" placeholder="Username" style="width: 100%; padding: 12px 20px; margin: 8px 0; box-sizing: border-box;">
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" placeholder="Password" style="width: 100%; padding: 12px 20px; margin: 8px 0; box-sizing: border-box;">
            <button id="login-btn" style="background-color: #337ab7; color: white; padding: 14px 20px; margin: 8px 0; border: none; cursor: pointer; width: 100%;">Login</button>
            <button id="close-modal" style="background-color: #f1f1f1; color: black; padding: 14px 20px; margin: 8px 0; border: 1px solid #ccc; cursor: pointer; width: 100%;">Cancel</button>
        </div>
        </div>

        <script>
        var modal = document.getElementById("login-modal");
        var openModal = document.querySelector("nav button");
        var closeModal = document.getElementById("close-modal");

        openModal.addEventListener("click", function () {
            modal.style.display = "block";
        });

        closeModal.addEventListener("click", function () {
            modal.style.display = "none";
        });

        window.addEventListener("click", function (event) {
            if (event.target == modal) {
                modal.style.display = "none";
            }
        });

        document.getElementById("login-btn").addEventListener("click", function () {
            var username = document.getElementById("username").value;
            var password = document.getElementById("password").value;
            // Implement your login functionality here
            console.log("Username:", username);
            console.log("Password:", password);
            modal.style.display = "none";
        });
        </script>
        </body>
        </html>
    `},
	"",
)
