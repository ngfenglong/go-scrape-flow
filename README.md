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
- **Integrated Logging:** Utilizes Logrus for logging, providing insights into the scraping process and aiding in identifying potential issues.
- **Data Export:** Supports exporting scraped data to Excel format, with flexibility in naming the output file using the `-o` flag.
- **Dynamic Content Extraction:** Currently enhancing the tool to allow users to specify custom selectors (like class names) for scraping specific content.
- **(In Progress) Single URL Analysis: Planning to extend functionality to analyze a single URL, extracting and listing all identifiable class names and types of data, and presenting them in an Excel file for comprehensive analysis.

These features are geared toward making `GoScrapeFlow` a versatile tool for web scraping, data analysis, and content aggregation, catering to both specific and broad scraping needs.

## Structure
```plaintext
├── cmd
│   ├── config.go
│   ├── root.go
│   └── start.go
├── helper
│   ├── excel
│   │   └── excel.go
│   └── log
│       └── log.go
├── httpclient
│   ├── http.go
│   └── proxy.go
├── output
│   [your output excel files will be here]
├── sitemap
│   ├── scrape.go
│   └── sitemap.go
├── main.go
├── Makefile
└── proxies.txt
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

### Quick Start with Makefile

To build and set up the project, run:

```sh
make all
```

This command builds the `goscrape` binary and sets up the necessary folders.

To run the project:

```sh
make run
```

For development testing with a predefined URL:

```sh
make dev-test
```

### Manual Setup and Run

If you prefer to set up and run without the Makefile:

1. **Build the project:**

    ```sh
    go build -o goscrape .
    chmod +x goscrape
    ```

2. **Create the output folder:**

    ```sh
    mkdir ./output
    ```

3. **Run the project:**

    ```sh
    ./goscrape
    ```

### Cleaning Up

To clean up the generated files:

- To remove the `goscrape` binary:

    ```sh
    make clean
    ```

- To remove the output folder:

    ```sh
    make clean-output
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


