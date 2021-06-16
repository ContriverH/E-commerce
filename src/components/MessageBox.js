import React from "react";

export default function MessageBox(props) {
  return (
    <div className={`alert alert-${props.varient || "info"}`}>
      {props.children}{" "}
      {/* the contents inside of the MessageBox will appear hear using children props*/}
    </div>
  );
}
