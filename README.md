# Gohyper - A (Go + HTMX) fullstack web app boilerplate

Gohyper is a fullstack web application boilerplate that embodies the principles of minimalism, enhanced developer experience, improved end-user experience, streamlined devops, code maintainability, adherence to better coding standards, and optimal code performance. It provides a solid foundation for building web applications using a carefully curated set of technologies.

## Technologies Used

1. **Golang for Backend**
   - A powerful and efficient programming language that facilitates rapid development of scalable and performant backend services.

2. **Chi as Fast Web Router**
   - A lightweight and expressive Go router that facilitates the creation of RESTful APIs and routing in web applications.

3. **Air for Watching Go Apps (in Development Mode)**
   - A live-reloading tool for Go applications that enhances the development experience by automatically rebuilding and restarting the application during development.

4. **Pollen CSS for Better Design Scales Guidelines**
   - A modular and customizable CSS framework that follows design scales to ensure consistency and aesthetics in web application design.

5. **HTMX for Web Interactivity and Reactivity**
   - A lightweight JavaScript library for creating web applications with minimal JavaScript while enhancing user interactivity and reactivity.

6. **Esbuild for JavaScript and CSS Bundling**
   - A fast JavaScript and CSS bundler for running, watching, and minifying JavaScript and CSS files.

7. **SQLite for Fast, Maintainable, and Embeddable Database**
   - A self-contained, serverless, and zero-configuration database engine that provides a reliable and efficient storage solution for web applications.

8. **SQLx for Database Access**
   - A database library for Go that provides a set of extensions on top of the standard database/sql library, making it easier to work with databases in Go.

9. **Makefile for Faster Execution of Commands**
   - A Makefile is included for faster and more convenient execution of common commands and tasks related to the development and deployment of the web application.

## Key Features

- **Minimalistic Structure**: Follows a minimalistic project structure for easier navigation and understanding.

- **Developer-Friendly**: Incorporates tools like Air and Esbuild for a seamless and efficient development experience.

- **User-Centric Design**: Utilizes HTMX and Pollen CSS to create an engaging and user-friendly web interface.

- **DevOps Streamlining**: Makefile included for faster and standardized execution of commands, improving DevOps experience.

- **Code Maintainability**: Follows best practices and coding standards to ensure maintainability and scalability.

- **Optimized Performance**: Leverages Golang's efficiency and other performance-oriented technologies to deliver fast and responsive web applications.

- **Minimal Dependencies**: Strives for minimal dependencies to keep the project lightweight and reduce potential issues.

- **Less to No Build Steps**: Minimizes the need for complex build steps, simplifying the development workflow.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/abhilash26/gohyper.git
   cd gohyper
   npm install
   go mod tidy
   make
