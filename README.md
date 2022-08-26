<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

Built With [![Go][Go]][React-url]


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://raw.githubusercontent.com/jerrychan807/imggg/master/image/20220826170908.png" alt="Logo" width="800" height="200">
  </a>

<h3 align="center">CakeSyrupPoolsLover</h3>

  <p align="center">
    Help PancakeLovers Earn More Usd 
    <br />
    薄饼糖浆池-新池提醒/查询-TGBot
    <br />
    <br />
    <a href="https://raw.githubusercontent.com/jerrychan807/imggg/master/image/20220822093506.png">View Demo</a>
    ·
    <a href="https://github.com/jerrychan807/cake-syrup-pools-lover/issues">Report Bug</a>
    ·
    <a href="https://github.com/jerrychan807/cake-syrup-pools-lover/issues">Request Feature</a>
  </p>
</div>


<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]][product-screenshot]

Blog: [Project-cake-syrup-pools-lover-PancakeSwap糖浆池TgBot-开发记录](https://jerrychan807.github.io/2022/20047.html)

<!-- GETTING STARTED -->
## Getting Started

### Use My Tg bot

1. https://t.me/J_bsc_bot
2. Click Menu Button

### Build Your Own Tg bot

1. Create your own bot from Key at [BotFather](https://t.me/BotFather)
2. Clone the repo
   ```sh
   git clone https://github.com/jerrychan807/cake-syrup-pools-lover.git
   ```
3. Install go packages
   ```sh
   go mod tidy
   ```
4. Enter your TgBot Token in `config.yaml` and Edit ConfigPath
   ```sh
   cp config/config.yaml.bak config/config.yaml
   ```
   ```sh
   vim config/config.yaml
   # Edit these lines
   # token :
   # projectFolder: 
   ```
   ```sh
   vim config/config.go
   # Edit this line
   viper.AddConfigPath("/jcoin/cake-syrup-pools-lover/config")
   ```
5. Init File and Folder
   ```sh
   echo "hello" > lastMd5.log
   mkdir tmp
   ```
6. Bulid and Run
   ```sh
   go build main.go
   ./main
   ```
   

<!-- USAGE EXAMPLES -->


<!-- ROADMAP -->
## Roadmap

- [ ] Input Address,query portfolio
- [ ] Calculate Apr  
- [ ] FeatureXXX



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/jerrychan807/cake-syrup-pools-lover.svg?style=for-the-badge
[contributors-url]: https://github.com/jerrychan807/cake-syrup-pools-lover/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jerrychan807/cake-syrup-pools-lover.svg?style=for-the-badge
[forks-url]: https://github.com/jerrychan807/cake-syrup-pools-lover/network/members
[stars-shield]: https://img.shields.io/github/stars/jerrychan807/cake-syrup-pools-lover.svg?style=for-the-badge
[stars-url]: https://github.com/jerrychan807/cake-syrup-pools-lover/stargazers
[issues-shield]: https://img.shields.io/github/issues/jerrychan807/cake-syrup-pools-lover.svg?style=for-the-badge
[issues-url]: https://github.com/jerrychan807/cake-syrup-pools-lover/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: https://raw.githubusercontent.com/jerrychan807/imggg/master/image/20220822093506.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[Go]: https://img.shields.io/badge/Go-v1.18.5-blue
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/