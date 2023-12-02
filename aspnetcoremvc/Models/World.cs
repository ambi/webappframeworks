using System.ComponentModel.DataAnnotations.Schema;

namespace aspnetcoremvc.Models;

[Table("world")]
public class World
{
    [Column("id")]
    public int Id { get; set; }

    // [IgnoreDataMember]
    // [NotMapped]
    // [JsonIgnore]
    // public int _Id { get; set; }

    [Column("randomnumber")]
    public int RandomNumber { get; set; }
}