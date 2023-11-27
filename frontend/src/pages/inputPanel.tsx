// React Component
import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { useState } from "react";

export default function InputPanel() {
	return (
		<div id="Input-Panel">
			<Tabs
				defaultActiveKey="Main"
				id="uncontrolled-tab-example"
				className="mb-3">
				<Tab eventKey="Main" title="Main">
					Tab content for Home
				</Tab>
				<Tab eventKey="Creator" title="Creator">
					Tab content for Profile
				</Tab>
				<Tab eventKey="Tags" title="Tags" disabled>
					Tab content for Contact
				</Tab>
			</Tabs>
		</div>
	);
}