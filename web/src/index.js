import React from "react";
import ReactDOM from "react-dom";
import "xwing-miniatures-font/dist/xwing-miniatures.css";

import "./index.css";

const xwingMan = [
  [],
  [0, 2, 2, 2],
  [1, 1, 2, 1, 1],
  [1, 1, 1, 1, 1],
  [0, 0, 1, 0, 0, 3]
];

const tieFMan = [
  [],
  [1, 0, 0, 0, 1],
  [2, 2, 2, 2, 2],
  [1, 2, 2, 2, 1, 0, 0, 0, 3, 3],
  [0, 0, 2, 0, 0, 3],
  [0, 0, 2]
];

class ManeuverCard extends React.Component {
  // renderEmptySquare creates the empty table-cells. Its key is denoted by the
  // column bearing.
  renderEmptySquare(bearing) {
    return <div className="empty-square" key={bearing} />;
  }

  // renderManeuverSquare returns the proper color and bearing to populate a
  // table-cell. Its key is denoted by the column bearing.
  renderManeuverSquare(bearing, difficulty) {
    const maneuver =
      difficulty + "-maneuvers-font xwing-miniatures-font-" + bearing;
    return (
      <div className="maneuver-square" key={bearing}>
        <i className={maneuver} />
      </div>
    );
  }

  // renderRow populates the row div with either empty-square or maneuver-square
  // objects.
  renderRow(row) {
    return row.map(maneuver => {
      const difficulty = maneuver[0];
      const bearing = maneuver[1];
      if (!difficulty) {
        return this.renderEmptySquare(bearing);
      }
      return this.renderManeuverSquare(bearing, difficulty);
    });
  }

  // renderSpeedSquare creates the speed square for a row. It is the first
  // element of a row div.
  renderSpeedSquare(speed) {
    return <div className="speed-square">{speed}</div>;
  }

  render() {
    const maneuvers = this.props.maneuvers;
    const rows = maneuvers.map((row, speed) => {
      speed = maneuvers.length - 1 - speed;
      return (
        <div className="maneuver-row" key={speed}>
          {this.renderSpeedSquare(speed)}
          {this.renderRow(row)}
        </div>
      );
    });

    return <div>{rows}</div>;
  }
}

class Ship extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isXwing: true
    };
  }

  switchShip() {
    this.setState({
      isXwing: !this.state.isXwing
    });
  }

  render() {
    const maneuvers = formatManeuvers(this.state.isXwing ? xwingMan : tieFMan);
    return (
      <div className="main">
        <div className="maneuver-card">
          <ManeuverCard maneuvers={maneuvers} />
        </div>
        <div className="side-button">
          <button onClick={() => this.switchShip()}>Switch Ship</button>
        </div>
      </div>
    );
  }
}

function formatManeuvers(maneuvers) {
  let longest = maneuvers.reduce((a, b) => (a.length > b.length ? a : b))
    .length;
  longest = longest > 6 ? longest : 6;
  let hasSloop = false;
  let hasTroll = false;

  let rows = maneuvers
    .map(row => {
      const fullRow = row.concat(Array(longest - row.length).fill(0));

      return fullRow.map((value, column) => {
        const difficulty = getDifficulty(value);
        const bearing = getBearing(column);
        if (difficulty) {
          switch (bearing) {
            case "sloopleft":
            case "sloopright":
              hasSloop = true;
              break;
            case "trollleft":
            case "trollright":
              hasTroll = true;
              break;
            default:
          }
        }

        return [difficulty, bearing];
      });
    })
    .reverse();
  let buffer = 0;

  if (hasSloop) {
    rows = rearangeManeuverRow(rows, buffer, 1);
    buffer += 2;
  }
  if (hasTroll) {
    rows = rearangeManeuverRow(rows, buffer, 3);
    buffer += 2;
  }

  return rows;
}

function rearangeManeuverRow(rows, buffer, offset) {
  return rows.map(row => {
    const end = buffer + 5;
    const left = end + offset - buffer;
    const right = end + offset - buffer + 1;

    return [row[left]].concat(
      row.slice(0, end),
      [row[right]],
      [row[end]],
      row.slice(right + 1)
    );
  });
}

function getBearing(bearing) {
  switch (bearing) {
    case 0:
      return "turnleft";
    case 1:
      return "bankleft";
    case 2:
      return "straight";
    case 3:
      return "bankright";
    case 4:
      return "turnright";
    case 5:
      return "kturn";
    case 6:
      return "sloopleft";
    case 7:
      return "sloopright";
    case 8:
      return "trollleft";
    case 9:
      return "trollright";
    default:
      return null;
  }
}

function getDifficulty(difficulty) {
  switch (difficulty) {
    case 1:
      return "white";
    case 2:
      return "green";
    case 3:
      return "red";
    default:
      return null;
  }
}
// =============================================================================

ReactDOM.render(<Ship />, document.getElementById("root"));
