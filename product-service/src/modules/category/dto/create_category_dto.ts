import { Transform } from 'class-transformer'
import { IsInt, IsMongoId, IsNotEmpty, IsOptional } from 'class-validator'
import { Types } from 'mongoose'
import { ToMongoId } from 'src/utils/mongodb'
import { TrimAllSpaces } from 'src/utils/trim_spaces'

export class CreateCategoryDto {
  @TrimAllSpaces()
  @IsNotEmpty()
  name: string

  @IsMongoId()
  @IsOptional()
  parentId: string
}

export class SaveCategoryDto {
  @TrimAllSpaces()
  @IsNotEmpty()
  name: string

  @ToMongoId()
  @IsMongoId()
  @IsOptional()
  parentId: Types.ObjectId = null

  @IsInt()
  @IsNotEmpty()
  right: number

  @IsInt()
  @IsNotEmpty()
  left: number
}
