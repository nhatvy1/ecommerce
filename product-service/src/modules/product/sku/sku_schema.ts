import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose'
import { HydratedDocument } from 'mongoose'

@Schema({ timestamps: true })
export class Sku {
  @Prop({ type: Array, default: [0] })
  sku_tier_idx: number[]
  /**
   * color = [red, green] = [0, 1]
   * size = [S, M] = [0, 1]
   * => red + S = [0, 0]
   * => red + M = [0, 1]
   */

  @Prop({ type: Boolean, default: false })
  default: boolean

  @Prop({ type: String, default: '' })
  slug: string

  @Prop({ type: Number, default: 0 })
  sort: number

  @Prop({ type: String, required: true })
  price: number

  @Prop({ type: Number, default: 0 })
  stock: number

  @Prop({ type: Boolean, default: true, index: true, select: false })
  isDraft: boolean

  @Prop({ type: Boolean, default: false, index: true, select: false })
  isPublished: boolean

  @Prop({ type: Boolean, default: false })
  isDelete: boolean
}

export type SkuDocument = HydratedDocument<Sku>
export const SkuSchema = SchemaFactory.createForClass(Sku)