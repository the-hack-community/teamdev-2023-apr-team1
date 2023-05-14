
export interface IStrayCat {
  catId: number;
  userId: string;
  photoData: string;
  captureDateTime: string;
  location: {
    lat: number;
    long: number;
  };
  name: string;
  features: string;
  condition: string;
  reactions: null;
}
