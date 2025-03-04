import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose'
import mongoose, { HydratedDocument } from 'mongoose'

@Schema({ timestamps: true })
export class Spu {
  @Prop({ type: String, required: true })
  name: string

  @Prop({ type: String, required: true })
  description: string

  @Prop({ type: String, default: '' })
  thumb: string

  @Prop({ type: String, default: '' })
  slug: string

  @Prop({ type: Number, required: true })
  price: number

  @Prop({ type: Boolean, default: true, index: true, select: false })
  isDraft: boolean

  @Prop({ type: Boolean, default: false, index: true, select: false })
  isPublished: boolean

  @Prop({ type: Boolean, default: false })
  isDeleted: boolean

  @Prop({ type: mongoose.Schema.Types.Mixed })
  product_attributes: any

  @Prop({ type: [mongoose.Schema.Types.Mixed], default: [] })
  product_variations: any[]
  /**
   * [
   *    {
   *      images: [],
   *      name: 'Color'.
   *      options: ['Red', 'Green']
   *    },
   *    { 
   *      images: []
   *      name: 'Size',
   *      options: ['S', 'M']
   *    }
   * ]
   */

  @Prop({
    type: Number,
    default: 4.5,
    min: [1, 'Rating must be above 1.0'],
    max: [5, 'Rating must be above 5.0']
  })
  product_ratingsAverage: number
}

export type SpuDocument = HydratedDocument<Spu> 
export const SpuSchema = SchemaFactory.createForClass(Spu)
