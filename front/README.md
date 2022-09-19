# 前端注意事项

+ 1.请在当天开始敲代码之前习惯性pull代码
>git pull
+ 2.请在写完代码后推送代码
>git push
+ 3.写完代码后请修改Readme.md的项目目录<br>好让大家伙知道目前前端整体进度如何



## 项目目录

```tree
.
├── README.md
├── index.html
├── package.json         # 依赖列表
├── public               # 公共资源文件夹，存放一些图片，字体等静态资源
│   └── vite.svg
├── src                  # 源代码文件夹，我们的代码都是在这个文件夹下写
│   ├── App.vue
│   ├── assets
│   │   └── vue.svg
│   ├── components
│   │ └── HelloWorld.vue
│   ├── main.ts
│   ├── style.css
│   └── vite-env.d.ts
├── tsconfig.json         # Typescript 的配置文件
├── tsconfig.node.json    # 同上
└── vite.config.ts        # 脚手架 vite 的配置文件