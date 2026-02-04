export interface UserSys {
  createdAt: Date;
  updatedAt: Date;
}

export interface User extends UserSys {
  id: number;
  username: string;
  email: string;
  role: string;
  isacutive: boolean;
}

export interface UserInfo {
  id: number;
  username: string;
  email: string;
  role: string;
}

export interface LoginUserInfo {
  username: string;
  password: string;
  email: string;
  role: string;
}
