import { IsOptional } from 'class-validator'
import { TrimAllSpaces } from 'src/utils/trim_spaces'

export class UpdateCategoryDto {
  @TrimAllSpaces()
  @IsOptional()
  name: string
}
