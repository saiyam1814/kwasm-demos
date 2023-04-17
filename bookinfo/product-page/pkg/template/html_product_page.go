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
                    background-color: #950eba;
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
                    background-color: #fff;
                    color: #222 !important;
                    border: none;
                    color: white;
                    float: right;
                    font-size: 14px;
                    margin-top: 8px;
                    padding: 5px 16px;
                    border-radius: 5px;
                }
                main {
                    display: flex;
                    justify-content: center;
                    padding: 20px;
                }
                h3 {
                    text-align: center;
                    color: #950eba;
                }
                p {
                    text-align: justify;
                }
                h4 {
                    text-align: center;
                    color: #950eba;
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
                .container {
                    padding-right: 15px;
                    padding-left: 15px;
                    margin-right: auto;
                    margin-left: auto;
                    width:1170px;
                    background-color:#950eba;
                }
                .container-fluid {
                    padding-right: 15px;
                    padding-left: 15px;
                    margin-right: auto;
                    margin-left: auto;
                    display: block;
                    width:100%;
                }
                .column1 {
                    width: 100%;
                    float: left;
                    position: relative;
                    min-height: 1px;
                    padding-right: 15px;
                    padding-left: 15px;
                    display: block;
                    text-align: center;
                }
                .row {
                    margin-right: -15px;
                    margin-left: -15px;
                }
                .column2 {
                    width: 50%;
                    float: left;
                    position: relative;
                    min-height: 1px;
                    padding-right: 15px;
                    padding-left: 15px;
                    display: block;
                    box-sizing: border-box;
                    max-width: 45vw;
                }
                .column3 {
                    box-sizing: border-box;
                    float: left;
                    position: relative;
                    min-height: 1px;
                    padding-right: 15px;
                    padding-left: 15px;
                    box-sizing: border-box;
                    display: block;
                    max-width: 45vw;
                }
                .served-by {
                    margin-left:40px;
                }
                dd {
                    margin-inline-start: 0px;
                }
                dt {
                    font-weight: 700;
                    line-height: 1.42857143;
                }
                blockquote {
                    font-size: 17.5px;
                    border-left: 5px solid #eee;
                    padding: 10px 20px;
                }
            </style>
        </head>
        <body>
        <nav>
            <div class="container">
                <a href="#">BookInfo Sample</a>
                <button type="button">Sign in</button>
            </div>
        </nav>
        <main>
            <div class="container-fluid">
                <div class="row">
                    <div class="column1">`,
		productPageSummaryReplaceTarget,
		`
                </div>
                <div class="row">
                    <div class="column2">
                        <h4>Book Details</h4>
        `,
		productPageDetailsReplaceTarget,
		`
                    </div>
                    <div class="column3">
                        <h4>Book Reviews</h4>
        `,
		productPageReviewsReplaceTarget,
		`
                    </div>
                </div>
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
