import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose'
import { HydratedDocument, Types } from 'mongoose'

@Schema({ timestamps: true })
export class Category {
  @Prop({ type: String, required: true })
  name: string

  @Prop({ type: Number, default: 0 })
  left: number

  @Prop({ type: Number, default: 0 })
  right: number

  @Prop({ type: Types.ObjectId, ref: 'Category' })
  parentId: Types.ObjectId
}

export type CategoryDocument = HydratedDocument<Category>

export const CategorySchema = SchemaFactory.createForClass(Category)
