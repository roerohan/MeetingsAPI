[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/roerohan/MeetingsAPI">
    <img src="https://project-logo.png" alt="Logo" width="80">
  </a> -->

  <h3 align="center">MeetingsAPI</h3>

  <p align="center">
    A simple API built using Go.
    <br />
    <a href="https://github.com/roerohan/MeetingsAPI"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/roerohan/MeetingsAPI">View Demo</a>
    ·
    <a href="https://github.com/roerohan/MeetingsAPI/issues">Report Bug</a>
    ·
    <a href="https://github.com/roerohan/MeetingsAPI/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contributors](#contributors-)



<!-- ABOUT THE PROJECT -->
## About The Project

A simple REST API is built for scheduling meetings as a starter project for Go. This has routes such as `/meetings` and `/meeting` to add meetings, and get meetings by their `id`. You can also filter meetings for a participant and by the start time and end time of the meeting.


### Built With

* [Go](https://golang.org/)
* [MongoDB](https://www.mongodb.com/)
* [Mongo-driver Go package](https://github.com/mongodb/mongo-go-driver)



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Go
- MongoDB

### Installation
 
1. Clone the Repo
```sh
git clone https://github.com/roerohan/MeetingsAPI.git
```
2. Install NPM packages
```sh
go get ./...
```



<!-- USAGE EXAMPLES -->
## Usage

To run the project, you can use:

```sh
go run src/main.go
```

1. POST `/meetings`: This route accepts a JSON of the following format:

```json
{
    "title": "something",
    "participants": [
        {
            "name": "something",
            "email": "something",
            "rsvp": "Yes"
        }
    ],
    "startTime": 1603059170289,
    "endTime": 1603059290289
}
```
> Note: This route ensures that a new meeting can't be added if a participant of the new meeting has already RSVP-ed "Yes" or "Maybe" to a different meeting. If it does, the list of meetings which conflict with the current meeting is returned.

2. GET `/meetings?participant=<email>`: This route takes the email ID of the participant and a list of meetings which the participant is included in, irrespective of the RSVP status.

3. GET `/meetings?start=<startTimestamp>&end=<endTimestamp>`: This route returns a list of all meetings that occur in the time duration between `start` and `end`.

4. GET `/meeting/<id>`: This route returns the meeting containing the ID as specified in the request parameter.


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/roerohan/MeetingsAPI/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

You are requested to follow the contribution guidelines specified in [CONTRIBUTING.md](./CONTRIBUTING.md) while contributing to the project :smile:.

<!-- LICENSE -->
## License

Distributed under the MIT License. See [`LICENSE`](./LICENSE) for more information.




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[roerohan-url]: https://roerohan.github.io
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=flat-square
[issues-url]: https://github.com/roerohan/MeetingsAPI/issues
