import React from "react";
// import PropTypes from "prop-types";

import XWingFont from "../Util/XWingFont";

function ManeuverCell(props) {
  return (
    <td className={"maneuver-cell " + props.cellType}>{props.children}</td>
  );
}

export function EmptyCell(props) {
  return <ManeuverCell cellType="empty-cell" />;
}

export function BearingCell(props) {
  const { bearing, difficulty } = props;
  return (
    <ManeuverCell
      cellType={"bordered-cell bearing-cell bearing-cell-" + difficulty}
    >
      <XWingFont fontType={"font"} symbol={bearing} />
    </ManeuverCell>
  );
}

export function SpeedCell(props) {
  return (
    <ManeuverCell cellType={"bordered-cell speed-cell"}>
      {props.speed}
    </ManeuverCell>
  );
}
