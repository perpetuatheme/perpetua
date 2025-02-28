<p align="center">
    <h1 align="center">Genesis</h1>
</p>

<p align="center">How the palettes were created</p>

## Inception

The original design of Perpetua arose from the desire to create a theme with
rich contrasts and beautiful natural colors that would work well in both dark
environments, such as a dark room at night, and bright environments, such as
during the day under the sun.

Every single color in the Perpetua palettes originates in some way from the
hues of the [NCS] colors. All non-monochromatic non-back colors have been
carefully selected to have an [APCA] contrast with the _Base 0_ color in the
same palette strictly greater than 60 and an [APCA] contrast with the
corresponding back color strictly greater than 48.

## Hues

The hues for _Red_, _Yellow_, _Green_, and _Blue_ are taken directly from [the
transformation] of the [NCS] colors to the [CIELAB color space]. In particular,
the mean values of the CIELAB hue angles ($h_{ab}$) are used, which correspond
to the following values:

| NCS hue triangle | Hue    |
| :--------------- | -----: |
| R                |  25.31 |
| Y                |  93.37 |
| G                | 164.93 |
| B                | 253.70 |

The hues for all other non-monochromatic colors are derived relative to these
four elementary colors. Specifically, by taking the hue values at the midpoint,
one-fourth, or three-fourths between the hues of two consecutive elementary
colors.

| Perpetua color | Point         | Hue value |
| :------------- | :------------ | --------: |
| Red            |     ―         |     25.31 |
| Orange         | Midpoint      |     59.34 |
| Yellow         |     ―         |     93.37 |
| Lime           | Midpoint      |    129.15 |
| Green          |     ―         |    164.93 |
| Turquoise      | One-fourth    |   ∼187.12 |
| Cyan           | Midpoint      |   ∼209.32 |
| Cerulean       | Three-fourths |   ∼231.51 |
| Blue           |     ―         |    253.70 |
| Violet         | One-fourth    |   ∼286.60 |
| Lavender       | Midpoint      |   ∼319.51 |
| Pink           | Three-fourths |   ∼352.41 |

The hues for the monochromatic colors are all the same within the same palette.
For the monochromatic colors of the _Light_ palette, the hue value has been
chosen to be one-fourth between _Orange_ and _Yellow_ to produce warm colors.
For the monochromatic colors of the _Dark_ palette, the hue value has been
chosen to be three-fourths between _Cerulean_ and _Blue_ to produce cool
colors. The hue values for the monochromatic colors of the _Light_ and _Dark_
palettes are approximately 180 degrees apart.

| Perpetua color             | Point         | Hue value |
| :------------------------- | :------------ | --------: |
|                            |               |           |
| Orange                     |     ―         |     59.34 |
| Light monochromatic colors | One-fourth    |    ~67.85 |
| Yellow                     |     ―         |     93.37 |
|                            |               |           |
| Cerulean                   |     ―         |   ∼231.51 |
| Dark monochromatic colors  | Three-fourths |   ~248.15 |
| Blue                       |     ―         |    253.70 |
|                            |               |           |

## Saturation and Lightness

The saturation and lightness values for the non-monochromatic colors have been
chosen to match the contrast expectations described in [Inception](#inception).
In addition, the saturation and lightness values for the back colors have been
specifically chosen to make them work well as background colors.

| Perpetua colors       | Saturation | Lightness |
| :-------------------- | ---------: | --------: |
| Light non-back colors |       99 % |      50 % |
| Light back colors     |       51 % |      88 % |
| Dark non-back colors  |      100 % |      76 % |
| Dark back colors      |       97 % |      34 % |

The saturation and lightness values for the monochromatic colors have been
chosen to meet the following three requirements:
- the highest [APCA] contrast of a background color with a text color must be
  between 80 and 85 to avoid eye strain caused by extreme contrast;
- the transition from the darkest to the lightest monochromatic color should be
  smooth, with only minor adjustments influenced by my personal taste;
- the _Base 0_, _Base 1_, and _Base 2_ colors should be closer to each other than the
  others to limit background variability.

| Perpetua color | Saturation | Saturation delta | Lightness | Lightness delta |
| :------------- | ---------: | ---------------: | --------: | --------------: |
|                |            |                  |           |                 |
| Light–Base 0   |       12 % |        ―         |      96 % |        ―        |
| Light–Base 1   |       11 % |             -1 % |      93 % |            -3 % |
| Light–Base 2   |       11 % |             -1 % |      90 % |            -3 % |
| Light–Base 3   |        9 % |             -1 % |      84 % |            -6 % |
| Light–Base 4   |        8 % |             -1 % |      78 % |            -6 % |
| Light–Base 5   |        7 % |             -1 % |      72 % |            -6 % |
| Light–Over 0   |        6 % |             -1 % |      66 % |            -6 % |
| Light–Over 1   |        5 % |             -1 % |      60 % |            -6 % |
| Light–Over 2   |        4 % |             -1 % |      54 % |            -6 % |
| Light–Text 2   |        3 % |             -1 % |      47 % |            -7 % |
| Light–Text 1   |        2 % |             -1 % |      40 % |            -7 % |
| Light–Text 0   |        1 % |             -1 % |      33 % |            -7 % |
|                |            |                  |           |                 |
| Dark–Base 0    |        9 % |        ―         |      10 % |        ―        |
| Dark–Base 1    |       10 % |             +1 % |      13 % |            +3 % |
| Dark–Base 2    |       11 % |             +1 % |      16 % |            +3 % |
| Dark–Base 3    |        9 % |             -2 % |      24 % |            +8 % |
| Dark–Base 4    |        8 % |             -1 % |      32 % |            +8 % |
| Dark–Base 5    |        7 % |             -1 % |      40 % |            +8 % |
| Dark–Over 0    |        6 % |             -1 % |      48 % |            +8 % |
| Dark–Over 1    |        5 % |             -1 % |      56 % |            +8 % |
| Dark–Over 2    |        4 % |             -1 % |      64 % |            +8 % |
| Dark–Text 2    |        3 % |             -1 % |      72 % |            +8 % |
| Dark–Text 1    |        2 % |             -1 % |      80 % |            +8 % |
| Dark–Text 0    |        1 % |             -1 % |      88 % |            +8 % |
|                |            |                  |           |                 |

## To-RGB and To-Hex Conversion

For color-picking, the [Okhsl color space] has been used, specifically to
convert hue, saturation, and lightness values to RGB (and hex) colors. The
[Okhsl color space] has been chosen to ensure consistency of lightness values
across different colors.

<!-- References -->

[NCS]: https://en.wikipedia.org/wiki/Natural_Color_System "Natural Color System - Wikipedia"
[APCA]: https://www.myndex.com/APCA/ "APCA (Accessible Perceptual Contrast Algorithm) Contrast Calculator"
[the transformation]: https://doi.org/10.1002/col.5080110211 "Derefeldt, G. and Sahlin, C. (1986), Transformation of NCS data into CIELAB colour space"
[CIELAB color space]: https://en.wikipedia.org/wiki/CIELAB_color_space "CIELAB color space - Wikipedia"
[Okhsl color space]: https://bottosson.github.io/posts/colorpicker/ "Two new color spaces for color picking - Okhsv and Okhsl"
