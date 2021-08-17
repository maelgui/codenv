import { IProxy } from "./proxy";

export interface IWorkspace {
    id: number;
    container_id: string;
	name:        string;
	image:       string;
	status:      string;
	proxies: Array<IProxy>;
}
