import React from "react";

import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
import { linkTo } from "@storybook/addon-links";

import ActionBar from "../components/ActionBar";
import Icon from "../components/Icon";
import Statline from "../components/Statline";
import ManeuverCard from "../components/ManeuverCard";
import XWingFont from "../components/Util/XWingFont";
import AltStyles from "../components/Util/AlternateStyles";

const xwingMan = {
  name: "X-wing",
  faction: ["Rebel Alliance"],
  attack: 3,
  agility: 2,
  hull: 3,
  shields: 2,
  actions: ["Focus", "Target Lock"],
  maneuvers: [
    [0, 0, 0, 0, 0, 0],
    [0, 2, 2, 2, 0, 0],
    [1, 1, 2, 1, 1, 0],
    [1, 1, 1, 1, 1, 0],
    [0, 0, 1, 0, 0, 3]
  ],
  size: "small",
  xws: "xwing",
  id: 0,
  firing_arcs: ["Front"]
};

// storiesOf("Welcome", module).add("to Storybook", () => (
//   <Welcome showApp={linkTo("Button")} />
// ));

// storiesOf("Button", module)
//   .add("with text", () => (
//     <Button onClick={action("clicked")}>Hello Button</Button>
//   ))
//   .add("with some emoji", () => (
//     <Button onClick={action("clicked")}>😀 😎 👍 💯</Button>
//   ));

storiesOf("ActionBar", module).add("default state", () => (
  <ActionBar actions={xwingMan.actions} />
));

storiesOf("Icon", module)
  .add("default ship", () => <Icon iconType={"ship"} symbol={xwingMan.xws} />)
  .add("default upgrade", () => (
    <Icon iconType={"upgrade"} symbol={"astromech"} />
  ));
storiesOf("Statline", module)
  .add("horizontal", () => (
    <Statline
      agility={xwingMan.agility}
      attack={xwingMan.attack}
      hull={xwingMan.hull}
      shields={xwingMan.shields}
    />
  ))
  .add("vertical", () => (
    <Statline
      altStyle={AltStyles.Vertical}
      agility={xwingMan.agility}
      attack={xwingMan.attack}
      hull={xwingMan.hull}
      shields={xwingMan.shields}
    />
  ));

storiesOf("ManeuverCard", module)
  .add("default state", () => <ManeuverCard maneuvers={xwingMan.maneuvers} />)
  .add("maxed state", () => (
    <ManeuverCard
      maneuvers={[
        [0, 0, 3],
        [0, 2, 2, 2, 0, 0, 0, 0, 0, 0, 3, 3, 3],
        [1, 1, 2, 1, 1, 0, 3, 3, 3, 3],
        [1, 1, 2, 1, 1, 3]
      ]}
    />
  ));
