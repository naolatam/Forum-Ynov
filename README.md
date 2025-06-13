# Warframe forum project
By Silk Road

## 📚 Description
This project is a forum based on the video game named Warframe, featuring moderation roles (as admin, moderator or a user, registered or not), live notification system and other elements which will be described in this file.

## ✨ Features
- ***Https protocol*** : the forum is using a Secured Socket Layout (SSL).
- ***Authentification page*** : you can be registered on the forum via OAuth login, using a Google/GitHub account.
- ***Post manager*** : on the forum you can manage your post by editing it or simply delete it.
- ***Pictures supporter*** : bring life to your post by adding pictures (JPEG, PNG, GIF).
- ***Moderation roles*** : you can be an admin, a moderator, a registered user or a simple guest. Those roles will determine your possibilities on the forum.
- ***Like/Dislike*** : as a registered user, you can like or dislike a post.
- ***Comment*** : as a registered user, you can comment a post.
- ***Live Notification system*** : as a registered user, you can receive notification, and have a 'recent activity' part displayed on your profile page which will show the last notifications received.

## 💻 Used Technologies
* ![TailwindCSS][tailwind-img]
* [![JS][JS-img]][JS-url]
* ![HTML5][HTML-IMG]
* ![CSS3][CSS-IMG]

## 🛠️ File Structure
- `src/`
- `pkg/`
- `repositories/`
- `services/`
- `dto/`

## 🌳 Tree Structure

```plaintext
.
├── docker-compose.yml
├── Dockerfile
├── init_db.sql
├── mariadb.env
├── src
│   ├── certs
│   │   ├── cert.pem
│   │   └── key.pem
│   ├── cmd
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── config
│   │   │   ├── database.go
│   │   │   └── env.go
│   │   ├── frontEnd
│   │   │   └── static
│   │   │       ├── css
│   │   │       │   ├── markdownVisualisation.css
│   │   │       │   ├── style.css
│   │   │       │   └── warframe-theme.css
│   │   │       ├── imgs
│   │   │       │   └── forums_background.jpg
│   │   │       └── js
│   │   │           ├── admin.js
│   │   │           ├── authentification.js
│   │   │           ├── editPublication.js
│   │   │           ├── findPublication.js
│   │   │           ├── globalFunction.js
│   │   │           ├── profile.js
│   │   │           └── publication.js
│   │   ├── handlers
│   │   │   ├── adminCategoryHandler.go
│   │   │   ├── adminContentHandler.go
│   │   │   ├── adminHandlers.go
│   │   │   ├── adminUserHandler.go
│   │   │   ├── authHandlers.go
│   │   │   ├── commentsHandler.go
│   │   │   ├── createPostHandler.go
│   │   │   ├── editPostHandler.go
│   │   │   ├── errorHandlers.go
│   │   │   ├── githubAuthHandler.go
│   │   │   ├── googleAuthHandler.go
│   │   │   ├── homeHandlers.go
│   │   │   ├── postReactionHandler.go
│   │   │   ├── postsHandlers.go
│   │   │   ├── profileHandlers.go
│   │   │   └── seePostHandler.go
│   │   ├── middleware
│   │   │   ├── helper.go
│   │   │   └── rateLimitMiddleware.go
│   │   ├── routes
│   │   │   ├── adminRoutes.go
│   │   │   ├── authRoutes.go
│   │   │   ├── errorRoutes.go
│   │   │   ├── homeRoutes.go
│   │   │   ├── postRoutes.go
│   │   │   ├── profileRoutes.go
│   │   │   └── routes.go
│   │   ├── server
│   │   │   └── server.go
│   │   └── templates
│   │       ├── admin.gohtml
│   │       ├── authentification.gohtml
│   │       ├── components
│   │       │   ├── footerComponent.gohtml
│   │       │   └── headerComponent.gohtml
│   │       ├── error.gohtml
│   │       ├── findPublication.gohtml
│   │       ├── index.gohtml
│   │       ├── profile.gohtml
│   │       ├── publicationEdit.gohtml
│   │       ├── publication.gohtml
│   │       └── template.go
│   └── pkg
│       ├── dtos
│       │   ├── githubUserInfo.go
│       │   ├── googleUserInfo.go
│       │   └── templates
│       │       ├── adminPageDto.go
│       │       ├── authenticationPageDto.go
│       │       ├── editPostPageDto.go
│       │       ├── errorPageDto.go
│       │       ├── headerDto.go
│       │       ├── homePageDtos.go
│       │       ├── postPageDto.go
│       │       ├── profilePageDto.go
│       │       ├── recentActivityDto.go
│       │       └── searchPostsDto.go
│       ├── hostedServices
│       │   ├── runner.go
│       │   └── sessionCleaner.go
│       ├── models
│       │   ├── category.go
│       │   ├── comment.go
│       │   ├── notification.go
│       │   ├── post.go
│       │   ├── reaction.go
│       │   ├── recentActivity.go
│       │   ├── report.go
│       │   ├── role.go
│       │   ├── session.go
│       │   └── user.go
│       ├── repositories
│       │   ├── categoryRepository.go
│       │   ├── commentRepository.go
│       │   ├── notificationRepository.go
│       │   ├── postRepository.go
│       │   ├── reactionRepository.go
│       │   ├── recentActivityRepository.go
│       │   ├── reportRepository.go
│       │   ├── repositories.go
│       │   ├── roleRepository.go
│       │   ├── sessionRepository.go
│       │   └── userRepository.go
│       ├── services
│       │   ├── categoryService.go
│       │   ├── commentService.go
│       │   ├── notificationService.go
│       │   ├── postService.go
│       │   ├── reactionService.go
│       │   ├── recentActivityService.go
│       │   ├── reportService.go
│       │   ├── roleService.go
│       │   ├── services.go
│       │   ├── sessionService.go
│       │   └── userService.go
│       └── utils
│           ├── auth.go
│           ├── certificate.go
│           ├── image.go
│           ├── oauth
│           │   ├── github.go
│           │   ├── google.go
│           │   └── oauthTools.go
│           └── timeAgo.go
└── tree.txt

```

## ▶️ How to Run
1) Clone the Repository on GitHub, using this link :
2) Navigate to the project directory : `src/cmd/main.go`
3) use the command : `go run main.go`
4) go on your web browser on this address : `https://localhost:8080`

## 👥 Contributors
This project was developed by:
<br>

<a href="https://github.com/DantesDels"><img src="https://avatars.githubusercontent.com/u/170110923?v=4" width="50" alt="Delver Sébastien"></a>
<a href="https://github.com/naolatam"><img src="https://avatars.githubusercontent.com/u/59016480?v=4" width="50" alt="Mata Loan"></a>
<a href="https://github.com/Torolgo"><img src="https://avatars.githubusercontent.com/u/190293274?v=4" width="50" alt="Dessenne Ylan"></a>


## 🔗 Links
- [![Trello][trello-img]](https://trello.com/b/qZNenZyC/forum-silk-road)
- [![Github][github-img]](https://github.com/naolatam/Forum-Ynov/)
- [![Canva][canva-img]](https://www.canva.com/design/DAGqDQIbMSQ/GlF-fkQFSRUjsb3sZLDRCw/edit?utm_content=DAGqDQIbMSQ&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)


[trello-img]: https://img.shields.io/badge/Trello-%23026AA7.svg?style=for-the-badge&logo=Trello&logoColor=white
[canva-img]: https://img.shields.io/badge/Canva-%2300C4CC.svg?style=for-the-badge&logo=Canva&logoColor=white
[gitea-img]: https://img.shields.io/badge/Gitea-%2300ACD7.svg?style=for-the-badge&logo=Gitea&logoColor=white
[github-img]: https://img.shields.io/badge/GitHub-%23121011.svg?style=for-the-badge&logo=github&logoColor=white

[tailwind-img]: https://img.shields.io/badge/TailwindCSS-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white
[JS-img]: https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E
[JS-url]: https://www.javascript.com
[HTML-IMG]: https://img.shields.io/badge/html5-%23E34F00.svg?style=for-the-badge&logo=html5&logoColor=white
[GO-IMG]: https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[CSS-IMG]: https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white