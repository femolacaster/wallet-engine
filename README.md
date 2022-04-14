<div align="center">
  <h1>Wallet Engine</h1>
  
  <p>
    A simple wallet engine microservices API implementation written in idiomatic GO.
  </p>
  

<br />

<!-- Table of Contents -->
# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  * [Screenshots](#camera-screenshots)
  * [Tech Stack](#space_invader-tech-stack)
  * [Environment Variables](#key-environment-variables)
- [Getting Started](#toolbox-getting-started)
  * [Prerequisites](#bangbang-prerequisites)
  * [Installation and Run](#gear-installation)
  * [Running Tests](#test_tube-running-tests)
  * [Deployment](#triangular_flag_on_post-deployment)
- [Usage](#eyes-usage)
- [Roadmap](#compass-roadmap)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

  

<!-- About the Project -->
## :star2: About the Project
A simple wallet engine microservices API implementation written in idiomatic GO.
Some technical considerations I took were to use a microservices implementation with a repository interface and an onion-style architecture around the domain-driven design. This is to make sure the Golang code is as idiomatic as possible, maintainable, reduces technical debts and release cycles, and can scale rapidly.

<!-- Screenshots -->
### :camera: Screenshots

<div align="center"> 
  <img src="https://placehold.co/600x400?text=Your+Screenshot+here" alt="screenshot" />
</div>


<!-- TechStack -->
### :space_invader: Tech Stack

<details>
  <summary>Client</summary>
  <ul>
    <li><a href="https://go.dev/">Golang</a></li>
  </ul>
</details>

<details>
<summary>Database</summary>
  <ul>
    <li><a href="https://www.mysql.com/">MySQL</a></li>
  </ul>
</details>

<details>
<summary>DevOps</summary>
  <ul>
    <li><a href="https://www.docker.com/">DirEnv</a></li>
  </ul>
</details>



<!-- Env Variables -->
### :key: Environment Variables

To run this project, you will need to add the following environment variables to your .envrc file

`DATABASE_URL`

<!-- Getting Started -->
## 	:toolbox: Getting Started

<!-- Prerequisites -->
### :bangbang: Prerequisites

This project uses go >= 1.17


<li><a href="https://go.dev/doc/install">Visit on how to install Golang</a></li>


<!-- Installation -->
### :gear: Installation and Run

Install my-project with npm

```bash
  yarn install my-project
  cd my-project
```
   
<!-- Running Tests -->
### :test_tube: Running Tests

To run tests, run the following command

```bash
  yarn test test
```

<!-- Usage -->
## :eyes: Usage

This is a generic wallet engine where you can generate a new wallet, add a debit/credit card, activate or deactivate a wallet.

<!-- Roadmap -->
## :compass: Roadmap

* [x] Add more tests
* [ ] Add authentication and multi-factor aunthentication
* [x] Add observability
* [ ] Add more microservices concept for durability



<!-- License -->
## :warning: License

Distributed under the MIT License. See LICENSE.txt for more information.


<!-- Contact -->
## :handshake: Contact

Olufemi Alabi - [@dev.to](https://dev.to/femolacaster) 


<!-- Acknowledgments -->
## :gem: Acknowledgements

Use this section to mention useful resources and libraries that you have used in your projects.


 - [Awesome README](https://github.com/Louis3797/awesome-readme-template/blob/main/README.md)
 - [Mario Carron's Golang Todo API microservice explanation](https://github.com/MarioCarrion/todo-api-microservice-example)
