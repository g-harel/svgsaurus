![svgsaurus](https://user-images.githubusercontent.com/9319710/43561245-2f4bb106-95e4-11e8-814e-25b2ddd87bff.png)

![λ](https://svgsaur.us?t=λ&h=40&w=24&y=29&s=28&c=999&o=b) ![svgsaurus](https://svgsaur.us?s=34&o=b&h=40&y=30&c=222) ![](https://svgsaur.us?t=&w=1&h=52)

<div style="text-align: center;">
      <img alt="purple" src="https://svgsaur.us?b=9400d3&t=_&h=2&w=148"
    /><img alt="blue" src="https://svgsaur.us?b=0000ff&t=_&h=2&w=148"
    /><img alt="green" src="https://svgsaur.us?b=00ff00&t=_&h=2&w=148"
    /><img alt="yellow" src="https://svgsaur.us?b=ffff00&t=_&h=2&w=148"
    /><img alt="orange" src="https://svgsaur.us?b=ff7f00&t=_&h=2&w=148"
    /><img alt="red" src="https://svgsaur.us?b=ff0000&t=_&h=2&w=148" />
</div>

&nbsp;

> svg generation by query string

* Live editing of output image using query params
* Fine-grained control over font's size, color, family, etc.
* Easily control padding and position of text
* Use background color to generate colorful accents
* Precisely position content with spacer images

## Usage ![](https://svgsaur.us?t=&w=1&h=40)

This project is deployed as a lambda function responding to requests at [svgsaur.us](https://svgsaur.us).

Use the url's query string to customize the generated svg image.

<table>
    <tr>
        <th>param</th>
        <th>description</th>
        <th>default</th>
        <th>example</th>
        <th width="200">result</th>
    </tr>
    <tr>
        <td><code>t</code></td>
        <td><code>text content</code></td>
        <td><code>svgsaurus</code></td>
        <td><code>https://svgsaur.us?t=abc123</code></td>
        <td><img height="40" src="https://svgsaur.us?t=abc123" /></td>
    </tr>
    <tr>
        <td><code>s</code></td>
        <td><code>font size</code></td>
        <td><code>55</code></td>
        <td><code>https://svgsaur.us?s=140</code></td>
        <td><img height="40" src="https://svgsaur.us?s=140" /></td>
    </tr>
    <tr>
        <td><code>h</code></td>
        <td><code>box height</code></td>
        <td><code>60</code></td>
        <td><code>https://svgsaur.us?h=30</code></td>
        <td><img src="https://svgsaur.us?h=30" /></td>
    </tr>
    <tr>
        <td><code>w</code></td>
        <td><code>box width</code></td>
        <td><code>265</code></td>
        <td><code>https://svgsaur.us?w=110</code></td>
        <td><img height="40" src="https://svgsaur.us?w=110" /></td>
    </tr>
    <tr>
        <td><code>x</code></td>
        <td><code>horizontal offset</code></td>
        <td><code>5</code></td>
        <td><code>https://svgsaur.us?x=40</code></td>
        <td><img src="https://svgsaur.us?x=40" /></td>
    </tr>
    <tr>
        <td><code>y</code></td>
        <td><code>vertical offset</code></td>
        <td><code>46</code></td>
        <td><code>https://svgsaur.us?y=10</code></td>
        <td><img src="https://svgsaur.us?y=10" /></td>
    </tr>
    <tr>
        <td><code>f</code></td>
        <td><code>font family</code></td>
        <td><code>arial</code></td>
        <td><code>https://svgsaur.us?f=impact</code></td>
        <td><img src="https://svgsaur.us?f=impact" /></td>
    </tr>
    <tr>
        <td><code>c</code></td>
        <td><code>font color</code></td>
        <td><code>000000</code></td>
        <td><code>https://svgsaur.us?c=999</code></td>
        <td><img src="https://svgsaur.us?c=999" /></td>
    </tr>
    <tr>
        <td><code>b</code></td>
        <td><code>box color</code></td>
        <td><code>&nbsp;</code></td>
        <td><code>https://svgsaur.us?b=999</code></td>
        <td><img src="https://svgsaur.us?b=999" /></td>
    </tr>
    <tr>
        <td><code>o</code></td>
        <td><code>options</code></td>
        <td><code>&nbsp;</code></td>
        <td><code>https://svgsaur.us?o=biu</code></td>
        <td><img src="https://svgsaur.us?o=biu" /></td>
    </tr>
    <tr>
        <td><code>r</code></td>
        <td><code>replacement</code></td>
        <td><code>_</code></td>
        <td><code>https://svgsaur.us?r=s</code></td>
        <td><img src="https://svgsaur.us?r=s" /></td>
    </tr>
</table>

## License

[MIT](./LICENSE)

&nbsp;

![made with computer](https://svgsaur.us/?t=Made_with_💻_by_✋&o=b&b=f8f8f8&s=20&w=888&h=260&x=360&y=135)
