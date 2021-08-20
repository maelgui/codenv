import { IProxy } from "./proxy";

export interface IWorkspace {
    id: string;
    container_id: string;
	name:        string;
	image:       string;
	status:      string;
	proxies: Array<IProxy>;
}
