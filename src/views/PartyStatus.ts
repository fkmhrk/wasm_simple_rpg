import Ractive from "ractive";
import { numberFormat } from "../models/numberFormat";

const PartyStatus = Ractive.extend(<any>{
  template: `<div class="party-state">
  {{#party.characters}}
  <div class="mdc-card character">
    <p>{{name}}</p>
<pre>HP{{format(hp, 3)}}
MP{{format(mp, 3)}}
Lv{{format(level, 3)}}</pre>
  </div>
  {{/}}
</div>`,
  data: () => ({
    format: numberFormat,
  }),
  css: `.party-state{
    display: flex;
  }
  .character {
    padding: 8px;
    width: 64px;
  }
  .character pre, .character p {
    margin: 0;
    font-family: Consolas, 'Courier New', Courier, Monaco, monospace;
  }
  `,

  _click: function() {
    this.fire("click");
  },
});

export default PartyStatus;
