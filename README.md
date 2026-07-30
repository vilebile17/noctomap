# NoctoMap

![Noctomap demo](demo.gif)

## About

Noctomap is a [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) built
using [BubbleTea](https://github.com/charmbracelet/bubbletea) for [Noctalia](https://noctalia.dev/) users.
It maps a wallpaper with a theme and allows users to switch both **simultaneously**.

## Motivation

While Noctalia does have support for themes generated from a wallpaper, I found that
those generated themes are never quite as good as _hand-crafted_ ones. Noctalia doesn't
have a feature to make wallpaper-theme combos and switch theme together so I decided
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
  "wallpaper_path": "~/Pictures/Wallpapers/pic.jpg", // The path to your wallpaper. It doesn't have to be ~/Pictures/Wallpapers/___ btw
  "colour_scheme_source": "builtin", // The source of the colour scheme. It is either 'builtin' 'community' 'wallpaper' or 'custom'
  "colour_scheme_name": "Catppuccin" // The actual name of the colour scheme. Be careful about spelling and capitalisation!
  "theme_mode": "dark" // In essence, just light- or dark-mode. Can be 'auto' if you want it to change according to the day-night cycle
}
```

> [!NOTE]
> There is a small inconsistency in Noctalia's naming system. If want to _generate a theme
> based upon a wallpaper_ by setting `colour_scheme_source` to `wallpaper`, then the `colour_scheme_name`
> won't be exactly as you see it in your settings. Instead everything is lower-case and replace all
> spaces with a hyphen. E.g. `M3 Tonal Spot` becomes `m3-tonal-spot`

_I have provided some examples in `noctomap.example.json` if you'd like a reference_

### Hyprland keybinds

If you're on [HyprLand](https://wiki.hypr.land) and you want to be able to change themes with quick notice, you can use the
provided `noctomap.sh` command. It opens a new terminal window, places it in the center of your screen
and resizes it (like seen in the demo gif at the top of this README)

Be sure to replace the initial `kitty` in `noctomap.sh` with which ever terminal you use!

If you want to bind this to a keybind then add this to your _hyprland config files_:

```lua
hl.bind(mainMod .. " + T", hl.dsp.exec_cmd("~/.config/hypr/config/noctomap.sh"))
```

and of course, customise the command to fit your use case (change the key, change the path etc.) 

## Contributing

If you would like to contribute anything to this project please **fork** the repo, make your changes
and then open a **pull request** to the main branch.
