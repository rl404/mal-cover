# MAL-Cover

Simple API to generate image cover CSS for MyAnimeList list.

## Installation

```bash
docker run -p 34001:34001 rl404/mal-cover
```

## Endpoints

### `/{user}/{type}`

Will generate CSS according to MyAnimeList username and type. For example:

- `https://mal-cover.herokuapp.com/rl404/anime?style=...`
- `https://mal-cover.herokuapp.com/rl404/manga?style=...`

### `/auto`

Will get username and type by reading URL page that call this endpoint. Just put this endpoint in CSS file like a normal import CSS file. For example:

- `@import url(https://mal-cover.herokuapp.com/auto?style=...);`
- `@\import "https://mal-cover.herokuapp.com/auto?style=...";`

## Styling

**This is the most important part**. Both endpoint needs a `style` param. The `style` value depends on how your list show your anime/manga cover image.

For example.

Your list's image cover style is like this.

```css
.animetitle[href*='/37716/']:before{
    background-image: url(https://myanimelist.cdn-dena.com/images/anime/1889/93555.jpg)
}
```

Convert it by replacing anime/manga id to `{id}` and image URL to `{url}`.

```css
.animetitle[href*='/{id}/']:before{background-image:url({url})}
```

Encode it using [URL encode](https://www.urlencoder.org/).

```
.animetitle%5Bhref%2A%3D%27%2F%7Bid%7D%2F%27%5D%3Abefore%7Bbackground-image%3Aurl%28%7Burl%7D%29%7D
```

Then use it in endpoint as `style` param.

```
https://mal-cover.herokuapp.com/rl404/anime?style=.animetitle%5Bhref%2A%3D%27%2F%7Bid%7D%2F%27%5D%3Abefore%7Bbackground-image%3Aurl%28%7Burl%7D%29%7D
```

Good luck.