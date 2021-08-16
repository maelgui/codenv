import axios from "axios";
import React from "react";
import { useParams } from "react-router-dom";
import { Terminal } from 'xterm';

import 'xterm/css/xterm.css';
import './terminal.css';

type TerminalViewParams = {
	id: string,
}

const TerminalView = () => {
	const { id: containerID } = useParams<TerminalViewParams>();

	const [taskID, setTaskID] = React.useState<string>();

	const termRef = React.createRef<HTMLDivElement>();
	const term = new Terminal();
	let socket: WebSocket | null = null

	React.useEffect(() => {
		axios.get(`/api/workspaces/${containerID}/exec`).then((response) => {
			console.log(response.data.task_id);
			setTaskID(response.data.task_id);
		})
	}, [containerID]);

	React.useEffect(() => {

		if (termRef.current && taskID) {
      socket = new WebSocket(`wss://${window.location.host}/ws/${taskID}`);
			term.open(termRef.current);
			term.onData(function (data) {
				if (socket) {
				socket.send(data);
				}
			});

			socket.onmessage = function (e) {
				term.write(e.data);
			}
		}
	}, [taskID]);

	return (
		<div ref={termRef}></div>
	)
};

export default TerminalView;
