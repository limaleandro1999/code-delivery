import { Prop, Schema, raw, SchemaFactory } from '@nestjs/mongoose';

export interface Position {
  lat: number;
  lng: number;
}

export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  startPosition: Position;

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  endPosition: Position;
}

export const RouteSchema = SchemaFactory.createForClass(Route);
