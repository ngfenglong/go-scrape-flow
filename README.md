# 🔍 GoScrapeFlow

GoScrapeFlow is a Go-based tool designed to efficiently scrape websites through sitemaps, offering extensibility for future logging and data exporting features.

> 🚧 **GoScrapeFlow is still under active development.** Expect continuous enhancements and potential changes. Documentation will be refined as the project evolves.

## Table of Contents

- [Features](#features)
- [Structure](#structure)
- [Installation](#installation)
- [Proxy Configuration](#proxy-configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## 🚀 Features

- **Efficient Scraping:** Optimized for scraping through sitemaps, ensuring comprehensive and systematic data extraction.
- **Modular Architecture:** Composed of organized modules for HTTP handling, proxy rotation, sitemap parsing, and scraping tasks.
- **(Upcoming) Logging: Integrated logging will provide insights into the scraping process and help identify potential issues.
- **(Upcoming) Data Export: Easily export your scraped data to various formats for analysis or storage.

## Structure
```plaintext
── cmd
│ └── main.go
├── httpclient
│ ├── http.go
│ └── proxy.go
└── sitemap
│ ├── sitemap.go
│ └── scrape.go

```

## 📦 Installation

1. Clone the repo:
    ```sh
    git clone https://github.com/ngfenglong/go-scrape-flow.git
    ```
2. Navigate to the directory:
    ```sh
    cd go-scrape-flow
    ```
3. Install the required dependencies (if any):
    ```sh
    go get
    ```

## Proxy Configuration

By default, GoScrapeFlow is configured to use proxy addresses from `proxies.txt`. Here's how you can manage this:

- **Update the Proxy List**: Modify the `proxies.txt` file to include your proxy addresses, one per line.
  
- **Disabling Proxy Usage**: If you're not planning to use any proxy (for reasons like having a VPN or a fast non-proxy connection), you can disable it:
  1. Open the `httpclient/http.go` file.
  2. Locate the `GetRequest` function.
  3. Comment out or remove the line: `refreshClientWithProxy()`.

## 📖 Usage

For quick start, you can now use the Makefile:

```sh
make start
```

If you prefer to run without Makefile:

```sh
go run cmd/main.go
```


## 🤝 Contributing

GoScrapeFlow is an open-source project, and contributions are warmly welcomed! Whether it's bug fixes, feature additions, or even documentation improvements, all forms of help are appreciated.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/NewFeature`)
3. Commit your changes (`git commit -m 'Add some NewFeature'`)
4. Push to the branch (`git push origin feature/NewFeature`)
5. Open a pull request

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.

## 📞 Contact

For any inquiries or clarifications related to this project, please contact [zell_dev@hotmail.com](mailto:zell_dev@hotmail.com).


