# NoctoMap

![Noctomap demo](demo.gif)

## About

Noctomap is a [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) built
using [BubbleTea](https://github.com/charmbracelet/bubbletea) for [Noctalia](https://noctalia.dev/) users.
It maps a wallpaper with a theme and allows users to switch both **simultaneously**.

## Motivation

While Noctalia does have support for themes generated from a wallpaper, I found that
those generated themes are never quite as good as _hand-crafted_ ones. Noctalia doesn't
have a feature to map wallpapers with themes and switch theme together so I decided
to implement it myself.

## Quickstart

First ensure that you have [Golang](https://go.dev) installed

Then clone the repo, install the program and copy the _config file_ to your home directory

```bash
git clone https://github.com/vilebile17/noctomap
cd noctomap
go install .
cp ./noctomap.example.json ~/.noctomap.json
```

## Configuration

Once you've got the `.noctomap.json` file in your home directory, you'll need to configure it
with your actual wallpaper-theme combos. Open it with the text editor of your choice
and add as many objects as you need to the `json` list.

Each object will be of this format:

```js
{
  "wallpaper_path": "~/Pictures/Wallpapers/pic.jpg", // quite obvious, the path to your wallpaper. It doesn't have to be ~/Pictures/Wallpapers/___ btw
  "colour_scheme_source": "builtin", // the source of the colour scheme. It is either 'builtin' 'community' 'wallpaper' or 'custom'
  "colour_scheme_name": "Catppuccin" // the actual name of the colour scheme. Be careful about capitalisation!
}
```

_Note that I have provided some examples in `.noctomap.example.json` if you'd like a reference_

## Contributing

If you would like to contribute anything to this project please **fork** the repo, make you changes
and then open a **pull request** to the main branch.
