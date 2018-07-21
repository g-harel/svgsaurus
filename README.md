<div style="text-align: center;">
      <img alt="purple" src="https://svgsaur.us?b=9400d3&t=_&h=4&w=120"
    /><img alt="purple" src="https://svgsaur.us?b=4b0082&t=_&h=4&w=120"
    /><img alt="blue" src="https://svgsaur.us?b=0000ff&t=_&h=4&w=120"
    /><img alt="green" src="https://svgsaur.us?b=00ff00&t=_&h=4&w=120"
    /><img alt="yellow" src="https://svgsaur.us?b=ffff00&t=_&h=4&w=120"
    /><img alt="orange" src="https://svgsaur.us?b=ff7f00&t=_&h=4&w=120"
    /><img alt="red" src="https://svgsaur.us?b=ff0000&t=_&h=4&w=120" />
</div>

# ![λ](https://svgsaur.us?t=λ&h=20&w=26&y=23&s=30&c=999) svgsaurus

> svg generation by query string

Spice up your READMEs with custom font-size, color, padding and much more!

## Usage

This project is deployed as a lambda function responding to requests at https://svgsaur.us.

Use the url's query string to customize the generated svg image.

Don't be affraid to get creative with emojis and background colors!

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

![made with computer](https://svgsaur.us/?t=Made_with_💻_by_✋&o=b&b=f1f1f1&s=20&w=888&h=260&x=360&y=135&f=arial)
