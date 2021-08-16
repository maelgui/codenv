import React from "react";
import { useParams } from "react-router-dom";
import { Terminal } from 'xterm';

import 'xterm/css/xterm.css';
import './terminal.css';

type TerminalViewParams = {
	id: string,
}

const TerminalView = () => {
	const { id } = useParams<TerminalViewParams>();
	const termRef = React.createRef<HTMLDivElement>();
	const term = new Terminal();
	let socket: WebSocket | null = null

	React.useEffect(() => {
		socket = new WebSocket(`ws://localhost:8080/api/workspaces/${id}/exec`);

		if (termRef.current) {
			console.log("connect");
			term.open(termRef.current);
			term.onData(function (data) {
				if (socket) {
				socket.send(data);
				console.log(socket)
				}
			});

			socket.onmessage = function (e) {
				term.write(e.data);
			}
		}
	}, []);

	return (
		<div ref={termRef}></div>
	)
};

export default TerminalView;
