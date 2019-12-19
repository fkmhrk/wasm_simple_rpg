import Ractive from "ractive";

const PartyStatus = Ractive.extend(<any>{
  template: `<div class="party-state">
  <div class="mdc-card character">
    <p>ABC</p>
    <p>HP123</p>
    <p>MP123</p>
    <p>Lv 12</p>
  </div>
  <div class="mdc-card character">
    <p>ABC</p>
    <p>HP123</p>
    <p>MP123</p>
    <p>Lv 12</p>
  </div>
  <div class="mdc-card character">
    <p>ABC</p>
    <p>HP123</p>
    <p>MP123</p>
    <p>Lv 12</p>
  </div>
  <div class="mdc-card character">
    <p>ABC</p>
    <p>HP123</p>
    <p>MP123</p>
    <p>Lv 12</p>
  </div>
  <div class="mdc-card character">
    <p>ABC</p>
    <p>HP123</p>
    <p>MP123</p>
    <p>Lv 12</p>
  </div>        
</div>`,
  css: `.party-state{
    display: flex;
  }
  .character {
    padding: 8px;
    width: 64px;
  }
  .character p {
    margin: 0;
    font-family: Consolas, 'Courier New', Courier, Monaco, monospace;
  }
  `,

  _click: function() {
    this.fire("click");
  },
});

export default PartyStatus;
