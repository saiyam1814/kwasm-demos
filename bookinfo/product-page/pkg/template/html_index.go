package template

import "strings"

const (
	indexTableReplaceTarget = "{TABLE}"
)

var indexHTML = strings.Join(
	[]string{
		`
			<!DOCTYPE html>
			<html>
			<head>
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
					table {
						background-color:transparent
					}
					.table {
						width: 100%;
						max-width: 100%;
						margin-bottom:20px;
						padding:5px;
					}
					.table-border, th, td {
						border:1px solid #ddd;
						box-sizing:border-box;
					}
					td, th {
						display: table-cell;
						padding:5px;
					}
					th:hover {
						background-color: #ebebeb;
					}
					td:hover {
						background-color: #ebebeb;
					}
					th {
						text-align: left;
					}
					h3 {
						font-size: 24px;
					}
					h1, h2, h3, h4, h5, h6 {
						margin-top: 20px;
						margin-bottom: 10px;
						font-family: inherit;
						font-weight: 500;
						line-height: 1.1;
						color: inherit;
					}
				</style>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1">

			</head>
			<body>
				
			<p>
				<h3>Hello! This is a simple bookstore application consisting of three services as shown below</h3>
			</p>
		`,
		indexTableReplaceTarget,
		`
			<p>
				<h4>Click on one of the links below to auto generate a request to the backend as a real user or a tester
				</h4>
			</p>
			<p><a href="/productpage?u=normal">Normal user</a></p>
			<p><a href="/productpage?u=test">Test user</a></p>

			</body>
			</html>
		`,
	},
	"",
)
