[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:b148630b554a7487 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# LeetCodeSolutions

三个相互独立的 Go 模块，分别对应已解决的三道 LeetCode 题目（21、27、2129），每个模块都有自己的 go.mod，并配有一个用于演练该解法的小型 main 函数。

[![CI](https://github.com/anyingiit/LeetCodeSolutions/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/LeetCodeSolutions/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/LeetCodeSolutions)](LICENSE)

[报告问题](https://github.com/anyingiit/LeetCodeSolutions/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/LeetCodeSolutions/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

本仓库收集了个人对 LeetCode 单道题目的解答，每道题都各自独立成一个自包含的 Go 模块：`problem21/go/problem21.go` 合并两个升序链表，`problem27/go/problem27.go` 原地移除切片中的指定值，`problem2129/go/problem2129.go` 将标题中的单词首字母大写。每个解法都自带 `go.mod`、一个 `main.go` 入口，以及讲解思路的注释。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/LeetCodeSolutions/issues)。

## 开始使用

### 环境要求

- Go 1.22 或更高版本，这是每个模块的 `go.mod` 所声明的下限（例如 `problem21/go/go.mod`）

### 安装

本仓库没有根模块：先克隆仓库，再进入你想要的那道题目自己的目录中构建。

```sh
git clone https://github.com/anyingiit/LeetCodeSolutions.git
cd LeetCodeSolutions
(cd problem21/go && go build ./...)
(cd problem27/go && go build ./...)
(cd problem2129/go && go build ./...)
```

## 用法

在某道题目自己的模块目录下直接运行它的 `main.go`：

```sh
cd problem21/go && go run .
cd ../../problem27/go && go run .
cd ../../problem2129/go && go run .
```

每个模块也都带有自己的自动化测试，例如：

```sh
cd problem21/go && go test ./...
```

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/LeetCodeSolutions](https://github.com/anyingiit/LeetCodeSolutions)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
