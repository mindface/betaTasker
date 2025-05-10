
export interface Posts {
 id: Number;
 title: string;
 text: string;
 name: string;
 disc: string;
 imgPath: string;
}

export type State = {
 status:String,
 posts: Posts[],
 rePosts: Posts[],
 lastUpdated: number,
 searchText: string,
 searchCategpryText: string,
 loading: boolean,
}

export type SetState = {
 status:String,
 posts: Posts[],
 rePosts: Posts[],
 lastUpdated: number,
 searchText: string,
 searchCategpryText: string,
 loading: boolean,
}